#name:logan.li
#use:v1.0  
#date:2023-02-07
#sudo docker build -t gofarm:1.0 .
#sudo docker run --name gofarm -p 8081:8081 -d gofarm:1.0

FROM golang:latest
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE="on"
#设置工作目录
WORKDIR /go/src/
RUN mkdir -p /go/src/mygofarm
COPY . /go/src/mygofarm
RUN cd /go/src/mygofarm \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  gofarm .

FROM alpine:latest
WORKDIR /root/
RUN mkdir -p /root/Settings
# 从编译阶段复制文件
# 这里使用了阶段索引值，第一个阶段从0开始，如果使用阶段别名则需要写成 COPY --from=build_stage /go/src/app/webserver /
COPY  --from=0 /go/src/mygofarm/gofarm /root/
COPY  --from=0 /go/src/mygofarm/Settings/config.yaml /root/Settings/config.yaml
RUN rm -rf /go/src/mygofarm/gofarm
#暴露端口
EXPOSE 8081
#最终运行docker的命令
ENTRYPOINT  ["/root/gofarm"]