## nginx 模块安装


#####1. 从容停止nginx
	ps -ef |grep nginx

	kill -QUIT 主进程号

#####2. 下载nginx module 

	 wget https://github.com/openresty/echo-nginx-module/archive/v0.60.tar.gz

#####3. 查看nginx编译参数

	/opt/nginx/sbin/nginx -V

#####4. 添加新的编译参数重新编译

     ./configure --add-module=/usr/local/src/echo-nginx-module-0.60

	重新编译 make -j2

#####5. 覆盖  nginx可执行文件

	cp objs/nginx /opt/nginx/sbin/

#####6 验证



参考： [https://github.com/openresty/echo-nginx-module#status](https://github.com/openresty/echo-nginx-module#status)