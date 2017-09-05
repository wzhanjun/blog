## gearman 学习

#### 安装gearman 

[http://www.linuxidc.com/Linux/2016-03/129348.htm](http://www.linuxidc.com/Linux/2016-03/129348.htm)

#### 安装php扩展

[http://pecl.php.net/get/gearman-1.0.2.tgz](http://pecl.php.net/get/gearman-1.0.2.tgz)


1、如果出现错误，可以在configure时指定"--with-gearman="，让php找到适合的libgearman

2、如果出现"[php_gearman.lo] Error 1"错误, 可以将"php_gearman.loT" 拷贝一个 "php_gearman.lo"，然后再make

3、如果启动php的时候gearman.so无效，则gearman扩展版本要到1.1.0版本才行

php-7 : [https://github.com/wcgallego/pecl-gearman](https://github.com/wcgallego/pecl-gearman)


问题：

 1. error while loading shared libraries: libmysqlclient.so.20: cannot open shared object file: No such file or directory

	1. find / -name 查找
	2. /etc/ld.so.conf/local.conf 添加目录 /usr/local/webserver/mysql/lib
	3. sudo ldconfig


* 修改php.ini 添加gearman扩展 重启php-pfm
	
	php -m


#### gearman 监控

(echo status; sleep 0.1) | nc 127.0.0.1 4730

