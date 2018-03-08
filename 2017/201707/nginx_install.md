## nginx 安装

#### 安装准备

1. 安装make

	yum -y install gcc automake autoconf libtool make

2. 安装g++

	yum install gcc gcc-c++

3. 依赖库

		- pcre
		- zlib
		- openssl

4. 安装
	
	  ./configure --prefix=/path/to/install --with-http_ssl_module \
	  --with-pcre=/path/to/pcre/src \
	  --with-zlib=/path/to/zlib/src \
	  --with-openssl=/path/to/openssl/src

	  make && make install

5. 检测

	查看80端口

    netstat -ano |grep 80

6. 启动

	/path/to/install/sbin/nginx


[安装教程](http://www.nginx.cn/install)

#### nginx 自启动脚本

[https://github.com/JasonGiedymin/nginx-init-ubuntu/blob/master/nginx](https://github.com/JasonGiedymin/nginx-init-ubuntu/blob/master/nginx)

#### 第三方模块安装

[http://www.ttlsa.com/nginx/how-to-install-nginx-third-modules/](http://www.ttlsa.com/nginx/how-to-install-nginx-third-modules/)


## openresty

最近跟着 章亦春 的nginx教程练习了一下

[http://blog.sina.com.cn/s/blog_6d579ff40100wi7p.html](http://blog.sina.com.cn/s/blog_6d579ff40100wi7p.html)

[http://openresty.org/#eBooks](http://openresty.org/#eBooks)

第一节讲的的echo 指令

下载echo-nginx-module 模块 重新编译

/path/to/nginx -V 查看原来的nginx 编译参数

[https://github.com/openresty/echo-nginx-module#installation](https://github.com/openresty/echo-nginx-module#installation)

安装完后重启nginx 原来的提示错误已经正确，