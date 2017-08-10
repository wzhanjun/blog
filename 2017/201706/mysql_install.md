#
mysql安装

#### mysql

* [mysql官方文档](http://dev.mysql.com/doc/refman/5.6/en/binary-installation.html)
* [mysql源码编译安装](http://www.cnblogs.com/xiongpq/p/3384681.html)
* [mysql主从配置](http://bestvivi.com/2015/09/06/MySQL%E5%A4%8D%E5%88%B6%E4%BB%8B%E7%BB%8D%E5%8F%8A%E6%90%AD%E5%BB%BA/) 

##### 编译参数

```

	cmake \
	-DCMAKE_INSTALL_PREFIX=/path/to/mysql \
	-DMYSQL_DATADIR=/path/to/mysql/data \
	-DSYSCONFDIR=/etc \
	-DWITH_MYISAM_STORAGE_ENGINE=1 \
	-DWITH_INNOBASE_STORAGE_ENGINE=1 \
	-DWITH_MEMORY_STORAGE_ENGINE=1 \
	-DWITH_READLINE=1 \
	-DMYSQL_UNIX_ADDR=/var/lib/mysql/mysql.sock \
	-DMYSQL_TCP_PORT=3306 \
	-DENABLED_LOCAL_INFILE=1 \
	-DWITH_PARTITION_STORAGE_ENGINE=1 \
	-DEXTRA_CHARSETS=all \
	-DDEFAULT_CHARSET=utf8 \
	-DDEFAULT_COLLATION=utf8_general_ci
    -DDOWNLOAD_BOOST=1 \
    -DWITH_BOOST=/path/to/mysql/boost/boost_1_59_0

```

	备注： data目录和 /var/lib/mysql 目录所有者和所属组都改为mysql:mysql
	
#### 添加mysql 账号

	shell> groupadd mysql
	shell> useradd -r -g mysql -s /bin/false mysql

#### 初始化


```
	As of MySQL 5.7.6, use the server to initialize the data directory:

    shell> bin/mysqld --initialize --user=mysql

	Before MySQL 5.7.6, use mysql_install_db:

	shell> bin/mysql_install_db --user=mysql


```


* 初始化密码

开启skip-grant-tables

1. vim /etc/my.cnf 加入skip-grant-tables

2. /etc/init.d/mysqld --skip-grant-tables start

```
     // SET PASSWORD = PASSWORD('new_password');
	 update user set authentication_string=password('YOURPASSWORDHERE') where user='root';

```


#### 自启动

```
cp /path/to/mysql/support-files/mysql.server /etc/init.d/mysqld

chmod 755 /etc/init.d/mysqld

chkconfig mysqld on

或  

sudo update-rc.d mysql defaults

```


#### mysql 添加到环境变量

```
vim /etc/profile

在末尾添加
PATH=$PATH:/usr/local/webserver/php/bin
export PATH

save

source /etc/profile

```
#### 授权远程登录

	grant all privileges on *.* to root@'%' identified by 'password' with grant option;

	FLUSH PRIVILEGES;


#### other

* mysql开启3360端口

```
telnet 192.168.xxx.xxx 3306 //检查端口

/etc/sysconfig/iptables

-A RH-Firewall-1-INPUT -m state --state NEW -m tcp -p tcp --dport 3306 -j ACCEPT
```

* mysql主从切换

```
1.从服务器上执行 SHOW PROCESSLIST 查看从服务器已经处理了日志中的所有语句
(Has read all relaylogwaiting for the slave I/O thread to update it）

2.STOP SLAVE IO_THREAD (从)

3.从库置为主库 STOP SLAVE和RESET MASTER和RESET SLAVE （从）

4.主库置为从库 reset master; change master to ....(执行切换master语句) ；
start slave; show slave status \G；

```

