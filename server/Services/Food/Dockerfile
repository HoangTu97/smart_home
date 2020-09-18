FROM golang:1.15.2-alpine

RUN apk add --no-cache gcc musl-dev bash git openssh make \
    && go get -u github.com/swaggo/swag/cmd/swag

ENV GOPROXY https://goproxy.cn,direct
ADD . $GOPATH/src/github.com/golang/smartHome
WORKDIR $GOPATH/src/github.com/golang/smartHome

EXPOSE 8080
ENTRYPOINT ["make", "dev"]