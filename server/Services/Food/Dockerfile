FROM golang:1.14-rc-alpine

ENV GOPROXY https://goproxy.cn,direct
ADD . $GOPATH/src/github.com/golang/smartHome
WORKDIR $GOPATH/src/github.com/golang/smartHome
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./smartHome"]