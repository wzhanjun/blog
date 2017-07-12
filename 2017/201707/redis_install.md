## redis安装


下载地址（http://redis.io/download）

	 make prefix=/path/to/install install

#### 配置

   daemonize   yes  #redis进程是否以守护进程的方式运行,yes为是,no为否(不以守护进程的方式运行会占用一个终端)

   pidfile /var/run/redis.pid  #指定redis进程的PID文件存放位置

   启动通过配置文件方式启动

#### phpredis

- 下载地址[https://github.com/phpredis/phpredis/archive/php7.zip](https://github.com/phpredis/phpredis/archive/php7.zip)

- 编译

	cd /path/to/phpredis

	/path/to/php/bin/phpize

	./configure --with-php-config=/path/to/php/bin/php-config

	make && make install

	修改php.ini
