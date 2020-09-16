FROM golang:1.15.2

RUN apt-get update && apt-get install -y build-essential && go get -u github.com/swaggo/swag/cmd/swag

ENV GOPROXY https://goproxy.cn,direct
ADD . $GOPATH/src/github.com/golang/smartHome
WORKDIR $GOPATH/src/github.com/golang/smartHome

EXPOSE 8080
ENTRYPOINT ["make", "dev"]