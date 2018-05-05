FROM golang:1.8

# This Dockerfile purpose is to build plugin in the same folder it's located
# docker build -t kafka_events_builder . && docker run -v /`pwd`/../../main/plugins/kafka_events:/go/out --rm --name kafka_events_build kafka_events_builder

WORKDIR /go/src/kafka_events
RUN go get github.com/Shopify/sarama/tools/kafka-console-producer


CMD ["bash"]
