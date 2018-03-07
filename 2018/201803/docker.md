##docker

#### 安装

	# 安装docker
	$ curl -sSL https://get.docker.io/ubuntu/ | sudo sh

	# 把docker用户加入 sudo组  避免每次命令都输入sudo
	$ sudo usermod -a -G docker $USER

	# 重启docker
	$ sudo service docker restart

	# 退出重新登录
	$ logout

#### docker命令
	
	#列出本机的所有 image 文件
	docker image ls

	#删除 image 文件
	docker image rm 

	#每运行一次，就会新建一个容器
	docker container run

	#启动已经生成、已经停止运行的容器文件
	docker container start

	#启动已经生成、已经停止运行的容器文件
	docker container stop [containerID]

	#在本机的另一个终端窗口，查出容器的 ID
	docker container ls
	
	# 停止指定的容器运行 (容器停止后文件并不会消失)
	docker container kill [containerID]

	#删除指定的容器文件
	docker container rm [containerID]

	#命令用来查看 docker 容器的输出
	docker container logs [containerID]

	#用于进入一个正在运行的 docker 容器
	docker container exec -it [containerID] /bin/bash

	#从正在运行的 Docker 容器里面
	docker container cp [containID]:[/path/to/file] .

#### docker homestead

1. https://laravel-china.org/topics/1855/replacing-laravel-homestead-development-environment-with-docker
2. http://laradock.io/introduction

#### 参考：

- [http://www.ruanyifeng.com/blog/2018/02/docker-tutorial.html](http://www.ruanyifeng.com/blog/2018/02/docker-tutorial.html)

- [https://realguess.net/2015/02/11/install-docker-and-then-run-docker-commands-without-sudo/](https://realguess.net/2015/02/11/install-docker-and-then-run-docker-commands-without-sudo/)

- [http://laradock.io/introduction/#features](http://laradock.io/introduction/#features)