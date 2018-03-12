# linux 脚本中运行php命令提示php file not found

1.  在shell脚本开头指定php路径

		#!/path/to/php

2. 添加php软连接

    	ln /path/to/php /usr/bin/php 

3. 在脚本中代码

		#!/usr/bin/env bash
		PHP=`which php`
		$PHP /path/to/php/file.php

4. 使用php的全路径

	/path/to/php /path/to/you/file.php


参考：

[https://stackoverflow.com/questions/5506913/bash-script-to-run-php-script](https://stackoverflow.com/questions/5506913/bash-script-to-run-php-script)
