FROM golang:1.9 as builder
WORKDIR /go/src/github.com/lvzhihao/uchat4influxdb
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /usr/local/uchat4influxdb
COPY --from=builder /go/src/github.com/lvzhihao/uchat4influxdb/uchat4influxdb .
ENV PATH /usr/local/uchat4influxdb:$PATH
