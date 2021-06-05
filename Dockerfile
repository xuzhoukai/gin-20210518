# https://www.cnblogs.com/panwenbin-logs/p/8007348.html

# FROM：指定基础镜像，必须为第一个命令
FROM golang:1.14

# LABEL：用于为镜像添加元数据
LABEL version="1.0" description="这是一个Web服务器" by="IT笔录"

# ENV：设置环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release \
    PORT=8080

# WORKDIR：工作目录，类似于cd命令
WORKDIR /app

# ADD 将本地文件添加到容器中，tar类型文件会自动解压(网络压缩资源不会被解压)，可以访问网络资源，类似wget
# COPY：功能类似ADD，但是是不会自动解压文件，也不能访问网络资源
COPY . .

# Run 构建镜像时执行的命令
RUN go build -o app main.go

RUN mkdir -p /app/bin && mv app /app/bin/

# EXPOSE：指定于外界交互的端口
EXPOSE 8080

# CMD：构建容器后调用，也就是在容器启动时才进行调用。

# VOLUME：用于指定持久化目录

# ENTRYPOINT：配置容器，使其可执行化。配合CMD可省去"application"，只使用参数。
ENTRYPOINT ["./bin/app"]