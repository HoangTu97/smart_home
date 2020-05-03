FROM golang:1.14-rc-alpine

ADD ../src /go/src/github.com/golang/smartHome

RUN go install github.com/golang/smartHome

ENTRYPOINT /go/bin/smartHome

EXPOSE 8080