#!/usr/bin/env bash


docker network create esorcerer-net
docker stop kafka
docker stop zookeeper
docker stop some-mongo
docker rm kafka
docker rm zookeeper
docker rm some-mongo
docker run -d --name zookeeper --network esorcerer-net zookeeper:3.4
docker run -d --name kafka --network esorcerer-net --env ZOOKEEPER_IP=zookeeper ches/kafka
docker run --name some-mongo --network esorcerer-net -d mongo



