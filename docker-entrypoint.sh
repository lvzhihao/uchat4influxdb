#!/bin/sh

# default timezone
if [ ! -n "$TZ" ]; then
    export TZ="Asia/Shanghai"
fi

# set timezone
ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
echo $TZ > /etc/timezone 

# k8s config  switch
if [ -f "/usr/local/uchat4influxdb/config/.uchat4influxdb.yaml" ]; then
    ln -s  /usr/local/uchat4influxdb/config/.uchat4influxdb.yaml /usr/local/uchat4influxdb/.uchat4influxdb.yaml
fi

# apply config
echo "===start==="
cat /usr/local/uchat4influxdb/.uchat4influxdb.yaml
echo "====end===="

# run command
if [ ! -z "$1" ]; then
    /usr/local/uchat4influxdb/uchat4influxdb $@
fi
