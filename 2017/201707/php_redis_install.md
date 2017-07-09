## php redis扩展编译

#### 编译扩展

     wget https://github.com/phpredis/phpredis/archive/2.2.4.tar.gz

     cd phpredis-2.2.7                      # 进入 phpredis 目录

     /usr/local/php/bin/phpize              # php安装后的路径

     ./configure --with-php-config=/usr/local/php/bin/php-config
    
     make && make install


#### 修改php.ini文件

	vi /usr/local/php/lib/php.ini

	增加如下内容:
	extension_dir = "/usr/local/php/lib/php/extensions/no-debug-zts-20090626"
	
	extension=redis.so

    重启php-fpm