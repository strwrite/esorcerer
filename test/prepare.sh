#!/usr/bin/env bash


docker network create kafka-net
docker stop kafka
docker stop zookeeper
docker rm kafka
docker rm zookeeper
docker run -d --name zookeeper --network kafka-net zookeeper:3.4
docker run -d --name kafka --network kafka-net --env ZOOKEEPER_IP=zookeeper ches/kafka

