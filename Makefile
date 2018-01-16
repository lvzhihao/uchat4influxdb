OS := $(shell uname)

build: */*.go
	go build 

dev:
	DEBUG=true go run main.go message

message:
	./uchat4influxdb message

docker-build:
	sudo docker build -t edwinlll/uchat4influxdb:latest .

docker-push:
	sudo docker push edwinlll/uchat4influxdb:latest

docker-ccr:
	sudo docker tag edwinlll/uchat4influxdb:latest ccr.ccs.tencentyun.com/wdwd/uchat4influxdb:latest
	sudo docker push ccr.ccs.tencentyun.com/wdwd/uchat4influxdb:latest
	sudo docker rmi ccr.ccs.tencentyun.com/wdwd/uchat4influxdb:latest

docker-uhub:
	sudo docker tag edwinlll/uchat4influxdb:latest uhub.service.ucloud.cn/mmzs/uchat4influxdb:latest
	sudo docker push uhub.service.ucloud.cn/mmzs/uchat4influxdb:latest
	sudo docker rmi uhub.service.ucloud.cn/mmzs/uchat4influxdb:latest
