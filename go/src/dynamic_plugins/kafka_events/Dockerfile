FROM golang:1.8

# This Dockerfile purpose is to build plugin in the same folder it's located
# docker build -t kafka_events_builder . && docker run -v /`pwd`/../../main/plugins/kafka_events:/go/out --rm --name kafka_events_build kafka_events_builder

WORKDIR /go/src/kafka_events
RUN go get gopkg.in/Shopify/sarama.v1
RUN go get gopkg.in/yaml.v2

COPY . .
RUN ( git config --global user.email 'str.write@gmail.com' \
&& git config --global user.name 'Timur Sultanaev' \
&& git init \
&& git add . && git commit -m init && git rev-parse HEAD > git.rev ) \
&& echo "Building revision $(cat git.rev) with GOPATH=$GOPATH" \
&& go build \
-ldflags \
"-pluginpath=kafka_events_custom_path \
 -X kafka_events_custom_path.build_date=`date -u '+%Y-%m-%d_%I:%M:%S%p'` \
 -X kafka_events_custom_path.git_hash=`cat git.rev`" \
-buildmode=plugin
#
CMD ["cp", "/go/src/kafka_events/kafka_events.so", "/go/out/"]
