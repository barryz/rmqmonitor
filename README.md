# rmqmonitor

[![Go Report Card](https://goreportcard.com/badge/github.com/barryz/rmqmonitor)](https://goreportcard.com/report/github.com/barryz/rmqmonitor)
[![Build Status](https://travis-ci.org/barryz/run.svg?branch=master)](https://travis-ci.org/barryz/rmqmonitor)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/oklog/run/master/LICENSE)

rmqmonitor is an agent that used for [open-falcon](http://open-falcon.org/) to monitoring [RabbitMQ](https://www.rabbitmq.com/).

## Arch Requirement
Linux

## Build

```bash
$make build
```

## Agent launch

```bash
$/bin/bash control.sh start|stop|restart
```
It will create a temporary directory `var` in your current path.

## Metrics

***overview metrics***:

| key | tag | type | note |
|-----|-----|------|------|
|rabbitmq.overview.publishRate| |GAUGE|rate of message publishing|
|rabbitmq.overview.deliverRate| |GAUGE|rate of message delivering|
|rabbitmq.overview.redeliverRate| |GAUGE|rate of message re-delivering|
|rabbitmq.overview.ackRate| |GAUGE|rate of message acknowledging|
|rabbitmq.overview.msgsTotal| |GAUGE|total messages(sum of unack and ready|
|rabbitmq.overview.msgsReadyTotal| |GAUGE|ready messages(not deliver yet)|
|rabbitmq.overview.msgsUnackTotal| |GAUGE|un-acknowledged messages|
|rabbitmq.overview.publishTotal| |GAUGE|total messaees of publishing|
|rabbitmq.overview.deliverTotal| |GAUGE|total messaees of delivering|
|rabbitmq.overview.redeliverTotal| |GAUGE|total messaees of re-delivering|
|rabbitmq.overview.channlesTotal| |GAUGE|total channels|
|rabbitmq.overview.connectionsTotal| |GAUGE|total connections|
|rabbitmq.overview.consumersTotal| |GAUGE|total counsumers|
|rabbitmq.overview.queuesTotal| |GAUGE|total queues|
|rabbitmq.overview.exchangesTotal| |GAUGE|total exchanges|
|rabbitmq.overview.isAlive| |GAUGE|healthy status of cluster|
|rabbitmq.overview.isPartition| |GAUGE|partition status of cluster|
|rabbitmq.overview.memUsedPct| |GAUGE|memory usage percentage|
|rabbitmq.overview.fdUsedPct| |GAUGE|percentage of fd usage|
|rabbitmq.overview.erlProcsUsedPct| |GAUGE|percentage of erlang processes|
|rabbitmq.overview.socketUsedPct| |GAUGE|percentage of socket usage|
|rabbitmq.overview.statsDbEvent| |GAUGE|the events of queue produced by management database|
|rabbitmq.overview.ioReadawait| |GAUGE|io_read_avg_wait_time|
|rabbitmq.overview.ioWriteawait| |GAUGE|io_write_avg_wait_time|
|rabbitmq.overview.ioSyncawait| |GAUGE|io_sync_avg_wait_time|
|rabbitmq.overview.memConnreader| |GAUGE|memory usage for connections reader|
|rabbitmq.overview.memConnwriter| |GAUGE|memory usage for connections writer|
|rabbitmq.overview.memConnchannels| |GAUGE|memory usage for connections channels|
|rabbitmq.overview.memMgmtdb| |GAUGE|memory usage for management db)|
|rabbitmq.overview.memMnesia| |GAUGE|memory usage for mnesia database)|
|rabbitmq.overview.runQueue| |GAUGE|total run_queues of Erlang|
|rabbitmq.overview.getChannelCost | |GAUGE|latency for getting channels|
|rabbitmq.overview.memAlarm| |GAUGE|memory alarm triggered|
|rabbitmq.overview.diskAlarm| |GAUGE|disc alarm triggered|

***Queue Metrics***

| key | tag | type | note |
|-----|-----|------|------|
|rabbitmq.queue.publish|name=$queue-name,vhost=$vhost|GAUGE|rate of message publishing with specified queue|
|rabbitmq.queue.delver_get|name=$queue-name,vhost=$vhost|GAUGE|rate of message delivering with specified queue|
|rabbitmq.queue.redeliver|name=$queue-name,vhost=$vhost|GAUGE|rate of message re-delivering with specified queue|
|rabbitmq.queue.ack|name=$queue-name,vhost=$vhost|GAUGE|rate of message acknowledging with specified queue|
|rabbitmq.queue.consumers|name=$queue-name,vhost=$vhost|GAUGE|total consumers with specified queue|
|rabbitmq.queue.consumer_utilisation|name=$queue-name,vhost=$vhost|GAUGE|consumers utilisation of the specified queue|
|rabbitmq.queue.dpratio|name=$queue-name,vhost=$vhost|GAUGE|the radio of deliver and publish with specified queue|
|rabbitmq.queue.memory|name=$queue-name,vhost=$vhost|GAUGE|total memories occupied with specified queue|
|rabbitmq.queue.messages|name=$queue-name,vhost=$vhost|GAUGE|total messages with specified queue|
|rabbitmq.queue.messages_ready|name=$queue-name,vhost=$vhost|GAUGE|ready messages with specified queue|
|rabbitmq.queue.messages_unacked|name=$queue-name,vhost=$vhost|GAUGE|un-ack messageis with specified queue|
|rabbitmq.queue.messages_status|name=$queue-name,vhost=$vhost|GAUGE|the status of the specified queue|


***Exchange Metrics***

| key | tag | type | note |
|-----|-----|------|------|
|rabbitmq.exchange.publish_in|name=$exchange-name,vhost=$vhost|GAUGE|publishing-inboud rate of the specified exchange|
|rabbitmq.exchange.publish_out|name=$exchange-name,vhost=$vhost|GAUGE|publishing-outboud rate of the specified exchange|
|rabbitmq.exchange.confirm|name=$exchange-name,vhost=$vhost|GAUGE|acknowledging rate of the specified exchange|

---

## Witch
spiderQ will starting a web server to handle several instructions which to control RabbitMQ process state.

The web server listening on port 5671 by default, it enable basicauth, and handle client's requests.

***RabbitMQ process management(graceful)***

```bash
curl -u noadmin:ADMIN -XPUT -d '{"name":"is_alive"}' http://127.0.0.1:5671/api/app/actions

curl -u noadmin:ADMIN -XPUT -d '{"name":"start"}' http://127.0.0.1:5671/api/app/actions

curl -u noadmin:ADMIN -XPUT -d '{"name":"stop"}' http://127.0.0.1:5671/api/app/actions

curl -u noadmin:ADMIN -XPUT -d '{"name":"restart"}' http://127.0.0.1:5671/api/app/actions
```

***Stop RabbitMQ process forcibly***

```bash
curl -u noadmin:ADMIN -XGET http://127.0.0.1:5671/api/app/fstop
```

***Get the healthy status of single RabbitMQ node***

```bash
curl -u noadmin:ADMIN -XGET http://127.0.0.1:5671/api/stats
```

***Start/Stop/Restart RabbitMQ statistics management database***

```bash
curl -u noadmin:ADMIN -XPUT -d '{"name":"reset"}' http://127.0.0.1:5671/api/stats/actions

curl -u noadmin:ADMIN -XPUT -d '{"name":"crash"}' http://127.0.0.1:5671/api/stats/actions

curl -u noadmin:ADMIN -XPUT -d '{"name":"terminate"}' http://127.0.0.1:5671/api/stats/actions
```
