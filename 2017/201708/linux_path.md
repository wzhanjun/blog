## linux 环境变量

#### 环境变量位置以及顺序
	
	1. /etc/profile (对所有用户生效)
	2. ~/.bash_profile |   ~/.bash_login  |  ~/.profile 
	3. ~/.bashrc  (如果~/.bash_profile存在)
	4. /etc/bashrc
	5. ~/.bash_logout
	
	
#### 添加方式

1. 命令行中执行 export PATH=/path/to/you/dir (临时生效)
2. 讲path 追加到以上文件中，保存 退出后执行  source /path/to/file

		export PATH=/path/to/you/dir:$PATH

#### 查看环境变量

	echo $PATH

