FROM golang:1.15.2-alpine

RUN apk add --no-cache gcc musl-dev git make \
    && go get -u github.com/swaggo/swag/cmd/swag

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/golang/smartHome
ADD . .

RUN rm -r ./data

EXPOSE 8080
ENTRYPOINT ["make", "dev"]