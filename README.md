# Golang版本RabbitMQ监控agent
-----

RabbitMQ 状态数据采集脚本
-----------------------------------------------------------

## Requirement:
- os: Linux

-----------------------------------------------------------

## 原理
-----------------------------------------------------------
通过RabbitMQ REST API 获取MQ相关状态数据，然后整合固定数据类型，推送到不同的监控系统。

## 相关指标
------------------------------------------------
- **overview指标**

| key | tag | type | note |
|-----|-----|------|------|
|rabbitmq.overview.publishRate| |GAUGE|生产总速率|
|rabbitmq.overview.deliverRate| |GAUGE|消费总速率|
|rabbitmq.overview.redeliverRate| |GAUGE|重新投递总速率|
|rabbitmq.overview.ackRate| |GAUGE|消费者确认总速率|
|rabbitmq.overview.msgsTotal| |GAUGE|消息总数, 等于ready + unack|
|rabbitmq.overview.msgsReadyTotal| |GAUGE|消息堆积总数|
|rabbitmq.overview.msgsUnackTotal| |GAUGE|消费未确认消息总数|
|rabbitmq.overview.publishTotal| |GAUGE|生产消息总数|
|rabbitmq.overview.deliverTotal| |GAUGE|投递消息总数|
|rabbitmq.overview.redeliverTotal| |GAUGE|重新投递消息总数|
|rabbitmq.overview.channlesTotal| |GAUGE|Channel 总数|
|rabbitmq.overview.connectionsTotal| |GAUGE|Connection 总数|
|rabbitmq.overview.consumersTotal| |GAUGE|Counsumer总数|
|rabbitmq.overview.queuesTotal| |GAUGE|队列总数|
|rabbitmq.overview.exchangesTotal| |GAUGE|exchange 总数|
|rabbitmq.overview.isAlive| |GAUGE|MQ健康状态(通过生产/消费判断集群读写)|
|rabbitmq.overview.isPartition| |GAUGE|MQ集群网络分区状态|
|rabbitmq.overview.memUsedPct| |GAUGE|内存使用占比|
|rabbitmq.overview.fdUsedPct| |GAUGE|file desc使用占比|
|rabbitmq.overview.erlProcsUsedPct| |GAUGE|Erlang 进程使用占比|
|rabbitmq.overview.socketUsedPct| |GAUGE|socket使用占比|
|rabbitmq.overview.statsDbEvent| |GAUGE|状态统计数据库事件队列个数|
|rabbitmq.overview.ioReadawait| |GAUGE|io_read_avg_wait_time|
|rabbitmq.overview.ioWriteawait| |GAUGE|io_write_avg_wait_time|
|rabbitmq.overview.ioSyncawait| |GAUGE|io_sync_avg_wait_time|
|rabbitmq.overview.memConnreader| |GAUGE|内存使用详情(connections reader)|
|rabbitmq.overview.memConnwriter| |GAUGE|内存使用详情(connections writer)|
|rabbitmq.overview.memConnchannels| |GAUGE|内存使用详情(connections channels)|
|rabbitmq.overview.memMgmtdb| |GAUGE|内存使用详情(management db)|
|rabbitmq.overview.memMnesia| |GAUGE|内存使用详情(Mnesia)|
|rabbitmq.overview.runQueue| |GAUGE|Erlang进程run_queue数量|

- **Queue指标**

| key | tag | type | note |
|-----|-----|------|------|
|rabbitmq.queue.publish|name=$queue-name,vhost=$vhost|GAUGE|该队列生产消息速率|
|rabbitmq.queue.delver_get|name=$queue-name,vhost=$vhost|GAUGE|该队列投递消息速率|
|rabbitmq.queue.redeliver|name=$queue-name,vhost=$vhost|GAUGE|该队列重新投递消息速率|
|rabbitmq.queue.ack|name=$queue-name,vhost=$vhost|GAUGE|该队列consumer确认消息速率|
|rabbitmq.queue.consumers|name=$queue-name,vhost=$vhost|GAUGE|该队列consumer个数|
|rabbitmq.queue.consumer_utilisation|name=$queue-name,vhost=$vhost|GAUGE|该队列消费利用率(消费能力)|
|rabbitmq.queue.dpratio|name=$queue-name,vhost=$vhost|GAUGE|该队列消费生产速率比|
|rabbitmq.queue.memory|name=$queue-name,vhost=$vhost|GAUGE|该队列所占内存字节数|
|rabbitmq.queue.messages|name=$queue-name,vhost=$vhost|GAUGE|该队列消息总数|
|rabbitmq.queue.messages_ready|name=$queue-name,vhost=$vhost|GAUGE|该队列等待被消费消息数|
|rabbitmq.queue.messages_unacked|name=$queue-name,vhost=$vhost|GAUGE|该队列消费未确认消息数|
|rabbitmq.queue.messages_status|name=$queue-name,vhost=$vhost|GAUGE|该队列状态(非idle/running,即认为不健康)|

