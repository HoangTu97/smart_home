FROM golang:1.15.2

RUN apt-get update
RUN apt-get install -y build-essential

ENV GOPROXY https://goproxy.cn,direct
ADD . $GOPATH/src/github.com/golang/smartHome
WORKDIR $GOPATH/src/github.com/golang/smartHome
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./smartHome"]