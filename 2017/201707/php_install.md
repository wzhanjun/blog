## 编译php


####php56编译

	./configure --prefix=/usr/local/php \
	--enable-fpm \
	--enable-mysqlnd \
	--enable-zip \
	--enable-mbstring \
	--with-pdo-mysql \
	--with-openssl \
	--with-mysql \
	--with-mysqli \
	--with-curl \
	--with-zlib \
	--with-gd \
	--with-mcrypt


#### fpm配置

    1. cp php-fpm.conf

    2. 修改配置
    groupadd www
    useradd -g www www

    修改www.conf  中的
	uesr=www  
	group=www

	3. 修改php-fpm通信方式
	listen =/var/run/php/php-fpm.sock;
	listen.owner = www
	listen.group = www
	listen.mode 0666

	4.设置时区
    设置时区  php.ini date.timezone ="UTC"


####php 安装后php-fpm自启动

	cp /path/to/source/sapi/fpm/init.d.php-fpm /etc/init.d/php-fpm
	
	sudo chmod +x /etc/init.d/php-fpm

	chkconfig --add php-fpm

	chkconfig php-fpm on

---


#### 缺少库

1. libxml2

	yum install libxml2
	
	yum install libxml2-devel -y


2. Cannot find OpenSSL's <evp.h>

	yum -y install openssl-devel
	
	find / -name "evp.h"

	--with-openssl-dir=/path/to/include/openssl

	ubuntu:

    sudo apt-get install libssl-dev libsslcommon2-dev


3. Please reinstall the libcurl distribution easy.h should be in /include/curl/

	yum -y install curl-devel


4. png.h not found.

	yum install libpng-devel

	sudo apt-get install libpng-dev


5. mcrypt.h not found. Please reinstall libmcrypt

	yum install libmcrypt libmcrypt-devel


##### 备注

	./configure --prefix=/usr/local/webserver/php71\
	--enable-fpm\
	--with-fpm-user=php-fpm\
	--with-fpm-group=www\
	--enable-phpdbg\
	--enable-mysqlnd\
	--with-mysqli=mysqlnd\
	--with-pdo-mysql=mysqlnd\
	--enable-opcache\
	--enable-pcntl\
	--enable-mbstring\
	--enable-soap\
	--enable-zip\
	--enable-calendar\
	--enable-bcmath\
	--enable-ftp\
	--enable-sockets\
	--with-openssl\
	--with-mhash\
	--with-zlib\
	--with-curl\
	--with-gettext\
	--with-gd\
	--enable-exif\
	--enable-gd-native-ttf\
	--enable-gd-jis-conv\
	--with-png-dir=/usr/lib\
	--with-jpeg-dir=/usr/lib\
	--with-freetype-dir=/usr/lib

	sudo make
	sudo make install
	sudo mkdir /etc/php
	sudo cp php.ini-development /etc/php/php.ini
	注意，以上PHP编译选项根据实际情况可调整。
	另外，还需要将PHP的可执行目录添加到环境变量中。 使用Vim/Sublime打开~/.bashrc，在末尾添加如下内容：
	export PATH=/usr/local/php/bin:$PATH
	export PATH=/usr/local/php/sbin:$PATH
	保存后，终端输入命令：
	source ~/.bashrc
