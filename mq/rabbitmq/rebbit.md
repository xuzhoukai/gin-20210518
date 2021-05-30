###centos7安装
https://www.cnblogs.com/fengyumeng/p/11133924.html

###启动停止服务
```shell script
#启动
rabbitmq-server -detached

#停止
rabbitmqctl stop

#查看状态
rabbitmqctl status
```

###配置用户
```shell script
#添加用户
rabbitmqctl add_user root root

#配置权限
rabbitmqctl set_permissions -p "/" root ".*" ".*" ".*"

#查看权限
rabbitmqctl list_user_permissions root

#删除用户
rabbitmqctl delete_user root

#设置用户为管理员，不然不能外部登录
rabbitmqctl set_user_tags root administrator

##查看用户列表
rabbitmqctl list_users
```
