## supervisor

#### 1. 安装

- centos

	yum info supervisor
	
	yum install supervisor

#### 2. 启动服务端

	supervisord -c /etc/supervisord.conf

	#检查是否运行
	ps aux|grep superviosrd

#### 3.新增项目配置

	配置详解：
	a)  在supervisord.conf文件中，分号“；”后面的内容表示注释
	b)  [group:组名]，设置一个服务分组，programs后面跟组内所有服务的名字，以分号分格。
	c)  [program：服务名]，下面是这个服务的具体设置：
	command:启用Tornado服务文件的命令，也就是我们手动启动的命令。
	directory:服务文件所在的目录
	user:启用服务的用户
	autostart:在 supervisord 启动的时候也自动启动
	autorestart:是否自动重启服务
	stdout_logfile：服务的产生的日起文件
	loglevel:日志级别
	startretries = 3; 启动失败自动重试次数，默认是 3
    startsecs = 5; 启动 5 秒后没有异常退出，就当作已经正常启动了

#### 4.客户端命令

	supervisorctl:

	status   		    # 查看程序状态
	stop    work-name   # 关闭 update_ip 程序
	start   work-name   # 启动 update_ip 程序
	restart work-name   # 重启 update_ip 程序
	reread  # 根据最新的配置文件,启动新配置或有改动的进程,配置没有改动的进程不会受影响而重启。
	update    ＃ 重启配置文件修改过的程序


#### 文档

- [官方文档](http://supervisord.org/)



