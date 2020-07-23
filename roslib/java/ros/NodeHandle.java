package com.roslib.ros;

import com.roslib.tinyros_msgs.TopicInfo;

import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.ThreadPoolExecutor;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

import com.roslib.tinyros_msgs.Log;

public class NodeHandle extends NodeHandleBase{
    private static final int MODE_FIRST_FF = 0;
    private static final int MODE_PROTOCOL_VER   = 1;
    private static final int MODE_SIZE_L         = 2;
    private static final int MODE_SIZE_L1        = 3;
    private static final int MODE_SIZE_H         = 4;
    private static final int MODE_SIZE_H1        = 5;
    private static final int MODE_SIZE_CHECKSUM  = 6;
    private static final int MODE_TOPIC_L        = 7;
    private static final int MODE_TOPIC_L1       = 8;
    private static final int MODE_TOPIC_H        = 9;
    private static final int MODE_TOPIC_H1       = 10;
    private static final int MODE_MESSAGE        = 11;
    private static final int MODE_MSG_CHECKSUM   = 12;
    private static java.lang.String TINYROS_LOG_TOPIC = "/tinyrosout";

    private HardwareTCP hardware_;
    private HardwareTCP loghd_;
    private boolean loghd_keepalive_;

    private boolean topic_list_recieved;
    private java.lang.String topic_list;

    private boolean service_list_recieved;
    private java.lang.String service_list;


    private byte[] message_in = new byte[INPUT_SIZE];
    private byte[] message_tmp = new byte[INPUT_SIZE];
    private byte[] message_out = new byte[OUTPUT_SIZE];


    private Publisher[] publishers = new Publisher[MAX_PUBLISHERS];
    private SubscriberT[] subscribers = new SubscriberT[MAX_SUBSCRIBERS];

    private Lock lock = new ReentrantLock();

    private ExecutorService spin_thread_pool_ = null;
    private ExecutorService spin_srv_thread_pool_ = null;
    private ExecutorService spin_log_thread_pool_ = null;
    private ExecutorService loghd_thread_pool_ = null;

    public NodeHandle() {
        for (int i = 0; i < MAX_PUBLISHERS; i++) {
            publishers[i] = null;
        }

        for (int i = 0; i < MAX_SUBSCRIBERS; i++) {
            subscribers[i] = null;
        }

        for (int i = 0; i < INPUT_SIZE; i++) {
            message_in[i] = 0;
            message_tmp[i] = 0;
        }

        for (int i = 0; i < OUTPUT_SIZE; i++) {
            message_out[i] = 0;
        }

        service_list = topic_list = "";

        hardware_ = new HardwareTCP();
        loghd_ = new HardwareTCP();
        ip_addr_ = node_name_ = "";
        loghd_keepalive_ = false;

        spin_ = true;

        spin_thread_pool_ = new ThreadPoolExecutor(3, 3, 0, TimeUnit.SECONDS,
                new ArrayBlockingQueue<>(1000), new ThreadPoolExecutor.DiscardOldestPolicy());
        spin_srv_thread_pool_ = new ThreadPoolExecutor(3, 3, 0, TimeUnit.SECONDS,
                new ArrayBlockingQueue<>(1000), new ThreadPoolExecutor.DiscardOldestPolicy());
        spin_log_thread_pool_ = new ThreadPoolExecutor(1, 1, 0, TimeUnit.SECONDS,
                new ArrayBlockingQueue<>(1000), new ThreadPoolExecutor.DiscardOldestPolicy());
        loghd_thread_pool_ = new ThreadPoolExecutor(1, 1, 0, TimeUnit.SECONDS,
                new ArrayBlockingQueue<>(1000), new ThreadPoolExecutor.DiscardOldestPolicy());
    }

    @Override
    public boolean initNode(java.lang.String node_name, java.lang.String ip_addr) {
        ip_addr_ = ip_addr;
		node_name_ = node_name;
        com.roslib.std_msgs.String msg = new com.roslib.std_msgs.String();
        if (hardware_.init(ip_addr_)) {
            msg.data = node_name_;
            publish(TopicInfo.ID_SESSION_ID, msg);
        }
        if(!loghd_keepalive_) {
            loghd_keepalive_ = true;
            loghd_thread_pool_.execute(new Runnable() {
                @Override
                public void run() {
                    try {
                        byte[] in = new byte[200];
                        while(loghd_keepalive_) {
                            if (!loghd_.connected()) {
                                if (loghd_.init(ip_addr_)) {
                                    com.roslib.std_msgs.String msg = new com.roslib.std_msgs.String();
                                    msg.data = node_name_ + "_log" ;
                                    publish(TopicInfo.ID_SESSION_ID, msg, true);
                                }
                                Thread.sleep(1000);
                            } else {
                                loghd_.read(in, in.length);
                            }
                        }
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                }
            });
        }
        return hardware_.connected();
    }

    @Override
    public int spin() {
        int mode = MODE_FIRST_FF;
        int rv = 0, bytes = 0, total_bytes = 0, index = 0;
        int checksum = 0, topic = 0, len = 1;

        lock.lock();
        for (int i = 0; i < MAX_PUBLISHERS; i++) {
            if (publishers[i] != null) {
                publishers[i].negotiated_ = false;
            }
        }
        for (int i = 0; i < MAX_SUBSCRIBERS; i++) {
            if (subscribers[i] != null) {
                subscribers[i].negotiated_ = false;
            }
        }
        lock.unlock();

        if (!hardware_.connected()) {
            return SPIN_ERR;
        } else {
            negotiateTopics();
        }

        while (spin_ && hardware_.connected()) {
            rv = hardware_.read(message_tmp, len);
            if (rv < 0) {
                mode = MODE_FIRST_FF;
                return SPIN_ERR;
            }

            for (int i = 0; i < rv; i++) {
                checksum += (message_tmp[i] & 0xff);
            }
            if (mode == MODE_MESSAGE) {
                for (int i = 0; i < rv; i++) {
                    message_in[index++] = message_tmp[i];
                    bytes--;
                }
                if (bytes == 0) {
                    len = 1;
                    mode = MODE_MSG_CHECKSUM;
                } else {
                    len = bytes;
                }
            } else if (mode == MODE_FIRST_FF) {
                if ((message_tmp[0] & 0xff) == 0xff) {
                    mode++;
                }
            } else if (mode == MODE_PROTOCOL_VER) {
                if ((message_tmp[0] & 0xff) == PROTOCOL_VER) {
                    mode++;
                } else {
                    mode = MODE_FIRST_FF;
                }
            } else if (mode == MODE_SIZE_L) {
                bytes = (message_tmp[0] & 0xff);
                index = 0;
                mode++;
                checksum = (message_tmp[0] & 0xff);
            } else if (mode == MODE_SIZE_L1) {
                bytes += (message_tmp[0] & 0xff) << 8;
                mode++;
            } else if (mode == MODE_SIZE_H) {
                bytes += (message_tmp[0] & 0xff) << 16;
                mode++;
            } else if (mode == MODE_SIZE_H1) {
                bytes += (message_tmp[0] & 0xff) << 24;
                total_bytes = bytes > 0 ? bytes : 1;
                mode++;
            } else if (mode == MODE_SIZE_CHECKSUM) {
                if ((checksum % 256) == 255)
                    mode++;
                else
                    mode = MODE_FIRST_FF;
            } else if (mode == MODE_TOPIC_L) {
                topic = (message_tmp[0] & 0xff);
                mode++;
                checksum = (message_tmp[0] & 0xff);
            } else if (mode == MODE_TOPIC_L1) {
                topic += (message_tmp[0] & 0xff) << 8;
                mode++;
            } else if (mode == MODE_TOPIC_H) {
                topic += (message_tmp[0] & 0xff) << 16;
                mode++;
            } else if (mode == MODE_TOPIC_H1) {
                topic += (message_tmp[0] & 0xff) << 24;
                mode = MODE_MESSAGE;
                if (bytes == 0)
                    mode = MODE_MSG_CHECKSUM;
                else
                    len = bytes;
            } else if (mode == MODE_MSG_CHECKSUM) {
                mode = MODE_FIRST_FF;
                if ((checksum % 256) == 255) {
                    if (topic == TopicInfo.ID_PUBLISHER) {
                        negotiateTopics();
                    } else if (topic == TopicInfo.ID_ROSTOPIC_REQUEST) {
                        com.roslib.std_msgs.String msg = new com.roslib.std_msgs.String();
                        msg.deserialize(message_in, 0);
                        topic_list = msg.data;
                        topic_list_recieved = true;
                    } else if (topic == TopicInfo.ID_ROSSERVICE_REQUEST) {
                        com.roslib.std_msgs.String msg = new com.roslib.std_msgs.String();
                        msg.deserialize(message_in, 0);
                        service_list = msg.data;
                        service_list_recieved = true;
                    } else if (topic == TopicInfo.ID_NEGOTIATED) {
                        TopicInfo ti = new TopicInfo();
                        ti.deserialize(message_in, 0);
                        for (int i = 0; i < MAX_PUBLISHERS; i++) {
                            if (publishers[i] != null && publishers[i].id_ == ti.topic_id) {
                                publishers[i].negotiated_ = ti.negotiated;
                            }
                        }
                        for (int i = 0; i < MAX_SUBSCRIBERS; i++) {
                            if (subscribers[i] != null && subscribers[i].id_ == ti.topic_id) {
                                subscribers[i].negotiated_ = ti.negotiated;
                            }
                        }
                    } else if (topic == TopicInfo.ID_TIME) {
                        syncTime(message_in);
                    } else {
                        SubscriberT sub = subscribers[topic - 100];
                        byte[] message = new byte[total_bytes];
                        System.arraycopy(message_in, 0, message, 0, total_bytes);
                        Runnable thread_pool_run = new Runnable() {
                            SubscriberT  sub_ = sub;
                            byte[] data = message;
                            @Override
                            public void run() {
                                if (sub_ != null) {
                                    sub_.callback(data);
                                }
                            }
                        };
                        if (subscribers[topic-100].topic_ == TINYROS_LOG_TOPIC) {
                            spin_log_thread_pool_.execute(thread_pool_run);
                        } else {
                            if (subscribers[topic-100].srv_flag_) {
                                spin_srv_thread_pool_.execute(thread_pool_run);
                            } else {
                                spin_thread_pool_.execute(thread_pool_run);
                            }
                        }
                    }
                }
            }
        }
        return SPIN_OK;
    }

    @Override
    public boolean ok() {
        return hardware_.connected();
    }

    public boolean advertise(Publisher p) {
        lock.lock();
        for (int i = 0; i < MAX_PUBLISHERS; i++) {
            if (publishers[i] == null) { // empty slot
                p.id_ = i + 100 + MAX_SUBSCRIBERS;
                p.nh_ = this;
                publishers[i] = p;
                lock.unlock();
                negotiateTopics(publishers[i]);
                return true;
            }
        }
        lock.unlock();
        return false;
    }

    public boolean subscribe(SubscriberT s) {
        lock.lock();
        for (int i = 0; i < MAX_SUBSCRIBERS; i++) {
            if (subscribers[i] == null) { // empty slot
                s.id_ = i + 100;
                subscribers[i] = s;
                lock.unlock();
                negotiateTopics(subscribers[i]);
                return true;
            }
        }
        lock.unlock();
        return false;
    }

    public boolean advertiseService(ServiceServer srv) {
        lock.lock();
        for (int i = 0; i < MAX_SUBSCRIBERS; i++) {
            if (subscribers[i] == null) { // empty slot
                srv.id_ = i + 100;
                subscribers[i] = srv;
                for (int j = 0; j < MAX_PUBLISHERS; j++) {
                    if (publishers[j] == null) { // empty slot
                        srv.pub.id_ = srv.id_;
                        srv.pub.nh_ = this;
                        publishers[j] = srv.pub;
                        lock.unlock();
                        negotiateTopics(subscribers[i]);
                        negotiateTopics(publishers[j]);
                        return true;
                    }
                }
            }
        }
        lock.unlock();
        return false;
    }

    public boolean serviceClient(ServiceClient srv) {
        lock.lock();
        for (int i = 0; i < MAX_SUBSCRIBERS; i++) {
            if (subscribers[i] == null) {// empty slot
                srv.id_ = i + 100;
                subscribers[i] = srv;
                for (int j = 0; j < MAX_PUBLISHERS; j++) {
                    if (publishers[j] == null) { // empty slot
                        srv.pub.id_ = srv.id_;
                        srv.pub.nh_ = this;
                        publishers[j] = srv.pub;
                        lock.unlock();
                        negotiateTopics(subscribers[i]);
                        negotiateTopics(publishers[j]);
                        return true;
                    }
                }
            }
        }
        lock.unlock();
        return false;
    }

    public void negotiateTopics()  {
        for (int i = 0; i < MAX_PUBLISHERS; i++) {
            if (publishers[i] != null) {
                negotiateTopics(publishers[i]);
            }
        }

        for (int i = 0; i < MAX_SUBSCRIBERS; i++) {
            if (subscribers[i] != null) {
                negotiateTopics(subscribers[i]);
            }
        }
    }

    @Override
    public int publish(long id, com.roslib.ros.Msg msg) {
        return publish(id, msg, false);
    }

    @Override
    public int publish(long id, com.roslib.ros.Msg msg, boolean islog) {
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
            if (!islog) {
                l = hardware_.write(message_out, l) ? l : -1;
            } else {
                l = loghd_.write(message_out, l) ? l : -1;
            }
            lock.unlock();
            return l;
        } else {
            lock.unlock();
            return -2;
        }
    }

    public java.lang.String getTopicList(int timeout) {
        com.roslib.std_msgs.String msg = new com.roslib.std_msgs.String();
        topic_list_recieved = false;
        topic_list = "";
        publish(TopicInfo.ID_ROSTOPIC_REQUEST, msg);
        long to = java.lang.System.currentTimeMillis() + timeout;
        while (!topic_list_recieved) {
            long now = java.lang.System.currentTimeMillis();
            if (now > to ) {
                System.out.println("Failed to get getTopicList: timeout expired");
                return "";
            }
            try {
                Thread.sleep(100);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
        return topic_list;
    }

    public java.lang.String getServiceList(int timeout) {
        com.roslib.std_msgs.String msg = new com.roslib.std_msgs.String();
        service_list_recieved = false;
        service_list = "";
        publish(TopicInfo.ID_ROSSERVICE_REQUEST, msg);
        long to = java.lang.System.currentTimeMillis() + timeout;
        while (!service_list_recieved) {
            long now = java.lang.System.currentTimeMillis();
            if (now > to) {
                System.out.println("Failed to get getServiceList: timeout expired\n");
                return "";
            }
            try {
                Thread.sleep(100);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
        return service_list;
    }

    public void log(int level, java.lang.String msg) {
        if (loghd_.connected()) {
            Log l = new Log();
            l.level = level;
            l.msg = "[" + node_name_ + "] " + msg;
            publish(TopicInfo.ID_LOG, l, true);
        }
    }

    @Override
    public void exit() {
        spin_ = false;
        loghd_keepalive_ = false;
        spin_thread_pool_.shutdown();
        spin_log_thread_pool_.shutdown();
        spin_srv_thread_pool_.shutdown();
        loghd_thread_pool_.shutdown();
        loghd_.close();
        hardware_.close();
    }

    @Override
    public void finalize() {
        exit();
    }
}
