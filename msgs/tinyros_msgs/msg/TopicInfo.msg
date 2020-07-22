# special topic_ids
uint32 ID_PUBLISHER=0
uint32 ID_SUBSCRIBER=1
uint32 ID_SERVICE_SERVER=2
uint32 ID_SERVICE_CLIENT=4
uint32 ID_ROSTOPIC_REQUEST=6
uint32 ID_ROSSERVICE_REQUEST=7
uint32 ID_LOG=8
uint32 ID_TIME=9
uint32 ID_NEGOTIATED=10
uint32 ID_SESSION_ID=11

# The endpoint ID for this topic
uint32 topic_id

string topic_name
string message_type

# MD5 checksum for this message type
string md5sum

# size of the buffer message must fit in
int32 buffer_size

bool negotiated

# node name
string node
