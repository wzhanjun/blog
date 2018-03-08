#linux 环境变量

#### 说明

1. /etc/profile

	此文件为系统的每个用户设置环境信息 

2. /etc/environment

 	设置的是整个系统的环境

3. ~/.bash_profile  

	 只对单个用户生效,当用户登录时该文件仅执行一次 (其他系统可能是： ~/.bash_profile ~/.bash_login)

4.  ~/.bashrc
 
	只对单个用户生效，当登录以及每次打开新的 shell 时，该文件被读取
	
5. export PATH=$PATH:/usr/local/test/bin

	方式只对当前终端 Shell 有效

6. source filepath 
	
	可以使变量设置在当前窗口立即生效