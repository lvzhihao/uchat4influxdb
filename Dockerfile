FROM golang:1.9 as builder
WORKDIR /go/src/github.com/lvzhihao/uchat4influxdb
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /go/src/github.com/lvzhihao/uchat4influxdb/uchat4influxdb .
