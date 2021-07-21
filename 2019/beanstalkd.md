# beanstalkd 安装使用

#### 项目地址

- [https://github.com/beanstalkd/beanstalkd](https://github.com/beanstalkd/beanstalkd)

- [https://beanstalkd.github.io/](https://beanstalkd.github.io/)

- [https://github.com/src-d/beanstool](https://github.com/src-d/beanstool)

#### 安装

- [下载地址](https://github.com/beanstalkd/beanstalkd/archive/v1.11.tar.gz)

	make

	make install PREFIX=/path/to/install

#### 添加日志持久化

	groupadd beanstalkd
	useradd -s /sbin/nologin -g beanstalkd beanstalkd
	mkdir -p /var/lib/beanstalkd/binlog
	chown -R beanstalkd:beanstalkd /var/lib/beanstalkd/binlog

#### 启动

	/usr/bin/beanstalkd -l 0.0.0.0 -p 11300 -b /var/lib/beanstalkd/binlog 


#### 参考

- [https://www.cnblogs.com/jiujuan/p/10887424.html](https://www.cnblogs.com/jiujuan/p/10887424.html)


