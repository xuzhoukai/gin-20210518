##centos7安装docker

https://www.cnblogs.com/yufeng218/p/8370670.html

###window10安装linux子系统
https://docs.microsoft.com/en-us/windows/wsl/install-win10#step-4---download-the-linux-kernel-update-package

###WslRegisterDistribution failed with error: 0x800701bc
https://blog.csdn.net/qq_44575789/article/details/107900485

###docker启动报错
基于WSL2 的 Docker Desktop 启动时 Failed to set version to docker-desktop: exit code: -1的解决方法

https://blog.csdn.net/qhd1994/article/details/111831427
```shell
netsh winsock reset
```

###创建镜像
```shell script
# 创建image
docker build -t gotest-image .
# 查看image
docker images
# 镜像创建container
docker run -itd --name gotest-image /bin/bash

# 设置镜像加速
"registry-mirrors": [
    "https://ebmfxi3g.mirror.aliyuncs.com"
  ]

# 
```