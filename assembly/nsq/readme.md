[nsqlookupd] 2021/08/14 18:28:39.127270 INFO: TCP: closing [::]:4160
localhost :: netcode/nsq-master/build » ./nsqlookupd
[nsqlookupd] 2021/08/14 18:28:43.030785 INFO: nsqlookupd v1.2.1-alpha (built w/go1.12.10)
[nsqlookupd] 2021/08/14 18:28:43.031286 INFO: TCP: listening on [::]:4160
[nsqlookupd] 2021/08/14 18:28:43.031308 INFO: HTTP: listening on [::]:4161


localhost :: netcode/nsq-master/build » ./nsqd --lookupd-tcp-address=127.0.0.1:4160
[nsqd] 2021/08/14 16:54:07.734379 INFO: nsqd v1.2.1-alpha (built w/go1.12.10)
[nsqd] 2021/08/14 16:54:07.734549 INFO: ID: 16
[nsqd] 2021/08/14 16:54:07.735162 INFO: TOPIC(test): created
[nsqd] 2021/08/14 16:54:07.735298 INFO: TOPIC(test): new channel(nsq_to_file)
[nsqd] 2021/08/14 16:54:07.735311 INFO: NSQ: persisting topic/channel metadata to nsqd.dat
[nsqd] 2021/08/14 16:54:07.745193 INFO: TCP: listening on [::]:4150
[nsqd] 2021/08/14 16:54:07.745224 INFO: HTTP: listening on [::]:4151
[nsqd] 2021/08/14 16:54:07.745262 INFO: LOOKUP(127.0.0.1:4160): adding peer
[nsqd] 2021/08/14 16:54:07.745296 INFO: LOOKUP connecting to 127.0.0.1:4160
[nsqd] 2021/08/14 16:54:07.746017 INFO: LOOKUPD(127.0.0.1:4160): peer info {TCPPort:4160 HTTPPort:4161 Version:1.2.1-alpha BroadcastAddress:localhost}
[nsqd] 2021/08/14 16:54:07.746041 INFO: LOOKUPD(127.0.0.1:4160): REGISTER test nsq_to_file
[nsqd] 2021/08/14 16:54:07.746207 INFO: LOOKUPD(127.0.0.1:4160): topic REGISTER test
[nsqd] 2021/08/14 16:54:07.746355 INFO: LOOKUPD(127.0.0.1:4160): channel REGISTER test nsq_to_file
[nsqd] 2021/08/14 16:55:01.285623 INFO: TCP: new client(127.0.0.1:49401)
[nsqd] 2021/08/14 16:55:01.285654 INFO: CLIENT(127.0.0.1:49401): desired protocol magic '  V2'


localhost :: netcode/nsq-master/build » ./nsqadmin --lookupd-http-address=127.0.0.1:4161
[nsqadmin] 2021/08/14 18:26:59.605035 INFO: nsqadmin v1.2.1-alpha (built w/go1.12.10)
[nsqadmin] 2021/08/14 18:26:59.605633 INFO: HTTP: listening on [::]:4171


localhost :: netcode/nsq-master/build » ./nsq_to_file --topic=my_topic_test --output-dir=/tmp --lookupd-http-address=127.0.0.1:4161
2021/08/14 18:29:34 INF    1 [my_topic_test/nsq_to_file] querying nsqlookupd http://127.0.0.1:4161/lookup?topic=my_topic_test
2021/08/14 18:29:34 INF    1 [my_topic_test/nsq_to_file] (localhost:4150) connecting to nsqd
[nsq_to_file] 2021/08/14 18:29:35.208730 INFO: [my_topic_test/nsq_to_file] opening /tmp/my_topic_test.localhost.2021-08-14_18.log
[nsq_to_file] 2021/08/14 18:29:35.208845 INFO: [my_topic_test/nsq_to_file] syncing 1 records to disk


