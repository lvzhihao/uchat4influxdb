FROM golang:1.9

WORKDIR /go/src/github.com/lvzhihao/uchat4influxdb

COPY . . 

RUN go-wrapper install && \
    rm -rf *
