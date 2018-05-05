#!/usr/bin/env bash


mkdir -p `pwd`/../../main/plugins/dynamic/kafka_events

docker build -t kafka_events_builder . && \
docker run \
-v /`pwd`/../../main/plugins/dynamic/kafka_events:/go/out \
--rm \
--name kafka_events_build kafka_events_builder && \
docker rmi kafka_events_builder
