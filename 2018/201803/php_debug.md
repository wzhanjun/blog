## php debug

#### 1. 安装xdebug扩展

	tar xzf xdebug-2.0.0RC4.tgz

	phpize

	./configure --php-config-path=

	sudo make && sudo make install


#### 2. phpStorm+homestead+Chrome Xdebug Helper

- 编辑xdebug.ini

    zend_extension=xdebug.so

	xdebug.remote_enable = 1

	xdebug.remote_connect_back = 1

	xdebug.remote_port = 9000

	xdebug.scream=0 

	xdebug.cli_color=1

	xdebug.show_local_vars=1

	重启php-fpm： sudo service php5-fpm restart


- 安装chrome扩展 Xdebug helper。

- 设置phpstorm ide key 为PHPStorm

- 设置目录映射:将项目目录和public目录映射到homestead对应的目录

- 运行 PHP Debug监听： Run>Start Listening for PHP Debug Connections。

- 把扩展调整为 debug 模式，如图
	
	
#### 3. 手动跟踪


	xdebug.trace_format = 0

	xdebug.auto_trace = On

	xdebug.trace_output_dir = /tmp/traces

	xdebug.trace_output_name = trace.%c.%p

	xdebug.collect_params = 4

	xdebug.collect_includes = On

	xdebug.collect_return = On

	xdebug.show_mem_delta = On


	在代码开始前调用：xdebug_start_trace() 

	在代码跟踪结束出调用： xdebug_stop_trace()

参考：

[https://laravel-china.org/topics/553/phpstorm-homestead-xdebug-chrome-helper-debug-configuration-xdebug](https://laravel-china.org/topics/553/phpstorm-homestead-xdebug-chrome-helper-debug-configuration-xdebug)

[https://www.ibm.com/developerworks/cn/opensource/os-php-xdebug/index.html#pests](https://www.ibm.com/developerworks/cn/opensource/os-php-xdebug/index.html#pests)

[https://www.zhihu.com/question/20348619](https://www.zhihu.com/question/20348619)