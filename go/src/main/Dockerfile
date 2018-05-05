FROM golang:1.8


WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...

RUN ( git config --global user.email 'str.write@gmail.com' \
&& git config --global user.name 'Timur Sultanaev' \
&& git init \
&& git add . && git commit -m init && git rev-parse HEAD > git.rev ) \
&& echo "Building revision $(cat git.rev) with GOPATH=$GOPATH" \
&& go install \
-ldflags \
"-X main.Buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.Githash=`cat git.rev`" \
-v ./...

CMD ["app"]
