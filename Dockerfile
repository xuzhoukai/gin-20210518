#源镜像
FROM golang:latest
#设置工作目录
WORKDIR $GOPATH/src/github.com/gotest
#将服务器的go工程代码加入到docker容器中
ADD . $GOPATH/src/github.com/gotest
#go构建可执行文件
RUN go build .
#暴露端口
EXPOSE 3000
#最终运行docker的命令
ENTRYPOINT ["./gotest"]