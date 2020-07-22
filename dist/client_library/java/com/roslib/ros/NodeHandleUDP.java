package com.roslib.ros;

import java.util.HashMap;
import java.util.Iterator;
import java.util.Set;
import java.util.UUID;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.ThreadPoolExecutor;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class NodeHandleUDP extends NodeHandleBase {
    private HardwareUDP hardware_;
    private Lock lock = new ReentrantLock();
    private byte[] message_in = new byte[INPUT_SIZE];
    private byte[] message_out = new byte[OUTPUT_SIZE];
    private boolean spin_;
    private ExecutorService spin_thread_pool_ = null;
    private ExecutorService negotiate_thread_pool_ = null;
    private boolean negotiate_keepalive_;

    private HashMap<Integer, Publisher> publishers_ = new HashMap<Integer, Publisher>();
    private HashMap<Integer, SubscriberT> subscribers_ = new HashMap<Integer, SubscriberT>();


    public NodeHandleUDP() {
        this.hardware_ = new HardwareUDP();
        for (int i = 0; i < INPUT_SIZE; i++) {
            message_in[i] = 0;
        }

        for (int i = 0; i < OUTPUT_SIZE; i++) {
            message_out[i] = 0;
        }
        this.spin_ = true;
        this.negotiate_keepalive_ = false;
        this.spin_thread_pool_ = new ThreadPoolExecutor(3, 3, 0, TimeUnit.SECONDS,
                new ArrayBlockingQueue<>(1000), new ThreadPoolExecutor.DiscardOldestPolicy());
        this.negotiate_thread_pool_ = new ThreadPoolExecutor(1, 1, 0, TimeUnit.SECONDS,
                new ArrayBlockingQueue<>(1000), new ThreadPoolExecutor.DiscardOldestPolicy());
    }

    private int generate_id() {
        UUID uuid = UUID.randomUUID();
        return uuid.hashCode();
    }

    @Override
    public boolean initNode(java.lang.String node_name, java.lang.String ip_addr) {
        node_name_ = node_name;
		ip_addr_ = ip_addr;
        if(!negotiate_keepalive_) {
            negotiate_keepalive_ = true;
            negotiate_thread_pool_.execute(new Runnable() {
                @Override
                public void run() {
                    try {
                        while(negotiate_keepalive_) {
                            negotiateTopics();
                            Thread.sleep(1000);
                        }
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                }
            });
        }
        return hardware_.init(ip_addr_);
    }

    @Override
    public int publish(long id, Msg msg, boolean islog) {
        lock.lock();
        int l = msg.serialize(message_out, 11);
        message_out[0] = (byte) 0xff;
        message_out[1] = (byte)(PROTOCOL_VER & 0xFF);
        message_out[2] = (byte)((l >> 0 ) & 0xFF);
        message_out[3] = (byte)((l >> 8 ) & 0xFF);
        message_out[4] = (byte)((l >> 16) & 0xFF);
        message_out[5] = (byte)((l >> 24) & 0xFF);
        message_out[6] = (byte)((255 - ((message_out[2] + message_out[3] + message_out[4] + message_out[5]) % 256)) & 0xFF);
        message_out[7] = (byte)((id >> 0 ) & 0xFF);
        message_out[8] = (byte)((id >> 8 ) & 0xFF);
        message_out[9] = (byte)((id >> 16) & 0xFF);
        message_out[10] = (byte)((id >> 24) & 0xFF);

        int chk = 0;
        for (int i = 7; i < l + 11; i++) {
            chk += message_out[i];
        }

        l += 11;

        message_out[l++] = (byte)((255 - (chk % 256)) & 0xFF);

        if (l <= OUTPUT_SIZE) {
            l = hardware_.write(message_out, l) ? l : -1;
            lock.unlock();
            return l;
        } else {
            lock.unlock();
            return -2;
        }
    }

    @Override
    public int publish(long id, Msg msg) {
        return publish(id, msg, false);
    }

    @Override
    public int spin() {
        while (spin_ && ok()) {
            if (subscribers_.size() <= 0) {
                try {
                    Thread.sleep(1000);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
                continue;
            }

            int rv = hardware_.read(message_in, INPUT_SIZE);
            if (INPUT_SIZE >= rv && rv > 0) {
                int topic = 0;
                int bytes = 0, index= 0, checksum = 0;
                do {
                    index = 0;
                    if ((message_in[index++] & 0xff) != 0xff) {
                        break;
                    }

                    if ((message_in[index++] & 0xff) != 0xb9) {
                        break;
                    }

                    bytes = (message_in[index] & 0xff);
                    bytes += (message_in[index + 1] & 0xff) << 8;
                    bytes += (message_in[index + 2] & 0xff) << 16;
                    bytes += (message_in[index + 3] & 0xff) << 24;
                    checksum = (message_in[index] & 0xff);
                    checksum += (message_in[index + 1] & 0xff);
                    checksum += (message_in[index + 2] & 0xff);
                    checksum += (message_in[index + 3] & 0xff);
                    checksum += (message_in[index + 4] & 0xff);
                    index += 5;

                    if((checksum % 256) != 255) {
                        break;
                    }

                    topic = (message_in[index] & 0x0ff);
                    topic += (message_in[index + 1] & 0xff) << 8;
                    topic += (message_in[index + 2] & 0xff) << 16;
                    topic += (message_in[index + 3] & 0xff) << 24;
                    checksum = (message_in[index] & 0xff);
                    checksum += (message_in[index + 1] & 0xff);
                    checksum += (message_in[index + 2] & 0xff);
                    checksum += (message_in[index + 3] & 0xff);
                    index += 4;

                    if (INPUT_SIZE < (index + bytes + 1)) {
                        break;
                    }

                    if(bytes > 0) {
                        for (int i=0; i < bytes + 1; i++) {
                            checksum += (message_in[index + i] & 0xff);
                        }
                    } else {
                        checksum += (message_in[index] & 0xff);
                    }

                    if ((checksum % 256) == 255) {
                        if(subscribers_.containsKey(topic)) {
                            SubscriberT  sub = subscribers_.get(topic);
                            int total_bytes = bytes > 0 ? bytes : 1;
                            byte[] message = new byte[total_bytes];
                            System.arraycopy(message_in, index, message, 0, total_bytes);
                            spin_thread_pool_.execute(new Runnable() {
                                SubscriberT  sub_ = sub;
                                byte[] data = message;
                                @Override
                                public void run() {
                                    if (sub_ != null) {
                                        sub_.callback(data);
                                    }
                                }
                            });
                        }

                    }
                } while(false);
            }
        }

        return SPIN_OK;
    }

    @Override
    public void exit() {
        spin_ = false;
        negotiate_keepalive_ = false;
        spin_thread_pool_.shutdown();
        negotiate_thread_pool_.shutdown();
        hardware_.close();
    }

    @Override
    public boolean ok() {
        return hardware_.connected();
    }

    public boolean advertise(Publisher p) {
        if (publishers_.size() >= MAX_PUBLISHERS) {
            return false;
        }
        p.id_ = generate_id();
        p.nh_ = this;
        p.negotiated_ = true;
        lock.lock();
        publishers_.put((int)p.id_, p);
        lock.unlock();
        negotiateTopics(p);
        return true;
    }

    public boolean subscribe(SubscriberT s) {
        if (subscribers_.size() >= MAX_SUBSCRIBERS) {
            return false;
        }
        s.id_ = generate_id();
        s.negotiated_ = true;
        lock.lock();
        subscribers_.put((int)s.id_, s);
        lock.unlock();
        negotiateTopics(s);
        return true;
    }

    public void negotiateTopics(){
        Set<Integer> keySet = publishers_.keySet();
        Iterator<Integer> it =keySet.iterator();
        while(it.hasNext()){
            negotiateTopics(publishers_.get(it.next()));
        }
        keySet = subscribers_.keySet();
        it =keySet.iterator();
        while(it.hasNext()){
            negotiateTopics(subscribers_.get(it.next()));
        }
    }
}
