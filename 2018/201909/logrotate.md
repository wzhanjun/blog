## logrotate 日志分割


#### 1. 配置

	/var/log/nginx/*.log {
		daily
		rotate 30
		missingok
		dateext
		compress
		delaycompress
		notifempty
		olddir oldpath
		su root root
		sharedscripts
		postrotate
    		[ -e /var/run/nginx.pid ] && kill -USR1 `cat /var/run/nginx.pid`
		endscript
	}


#### 2. 测试

	/usr/sbin/logrotate -d -f /etc/logrotate.conf

#### 3. 定时任务

	59 23 * * *  /usr/sbin/logrotate -f /etc/logrotate.d/nginx

参考：[https://www.jianshu.com/p/ea7c2363639c](https://www.jianshu.com/p/ea7c2363639c)






