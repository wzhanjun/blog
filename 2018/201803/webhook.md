## webhook

	服务器通过git来部署代码,本地提交代码后服务器自动拉取更新。

#### 新建git项目

	git clone xxxxx

#### 编写自动部署脚本

	<?php
	// 这里只是最简单的拉取代码，如果要做更加多的操作，如验证、日志，请自己解析push内容并操作
	
	// 获取push数据内容的方法
	$requestBody = file_get_contents("php://input");
	
	// 只需这一行代码便可拉取
	shell_exec("cd /var/www/Project && git pull"); // 目录换成项目的目录
	
	?>


#### 配置webhook
	
	在仓库找到webhooks设置项，

	url填写部署脚本地址：http://example.com/hook/index.php 


#### 配置权限

	部署脚本由服务器运行(假如www), 项目git部署用户(假如：root)可能不是同一个，执行部署脚本时会提示没权限.

	visudo
	
	www ALL = (www) NOPASSWD: /usr/bin/git

	The first ALL is the users allowed
	The second one is the hosts
	The third one is the user as you are running the command
	The last one is the commands allowed
	


参考：

[https://segmentfault.com/a/1190000012643992](https://segmentfault.com/a/1190000012643992)

[https://gist.github.com/cowboy/619858](https://gist.github.com/cowboy/619858)[leymannx]的回答

[https://unix.stackexchange.com/questions/201858/what-does-all-all-all-all-mean-in-sudoers](https://unix.stackexchange.com/questions/201858/what-does-all-all-all-all-mean-in-sudoers)

[https://segmentfault.com/a/1190000007394449](https://segmentfault.com/a/1190000007394449)
