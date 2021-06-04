FROM golang:1.14

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release \
    PORT=8080

WORKDIR /app

COPY . .

RUN go build -o app main.go

EXPOSE 8080 8080

ENTRYPOINT ["./app"]