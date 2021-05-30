###race使用
参考url：https://lailin.xyz/post/go-training-week3-data-race.html
```shell script
#添加 -race 开启，竞争检查
go run -race main.go

#环境配置配置race参数，遇到第一个数据竞争就退出
GORACE=halt_on_error=1
```