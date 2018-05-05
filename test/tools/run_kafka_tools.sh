#!/usr/bin/env bash

docker build -t kafka_tools . && \
docker run -it --rm --network esorcerer-net --name kafka_tools kafka_tools

#kafka-console-producer -topic=my_topic -value=value -brokers=kafka:9092
