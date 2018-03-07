#postgresql 安装

####创建用户

	#不建议-M 不然后面大堆警告
	useradd  postgres   
	#添加权限
	mkdir /usr/local/postgresql
	chown -R postgres:postgres /usr/local/postgresql
	chmod -R  755 /usr/local/postgresql


#### 添加环境变量

	cat > ~/.bashrc <<"EOF"
	export PATH=$PATH:/usr/local/postgresql/bin
	#启动数据库目录，配置后启动的时候会省掉-D选项
	export PGDATA=/usr/local/postgresql/data
	EOF
	#使之生效
	source ~/.bashrc

#### 安装依赖
	
	yum -y install gcc systemtap-sdt-devel.x86_64  perl-ExtUtils-Embed readline-devel zlib-devel openssl-devel   pam-devel libxml2-devel  	libxslt-devel  openldap-devel  tcl-devel  python-devel uuid-devel 

####安装

	su - postgres
	
	cd  /usr/local/src/postgresql-10beta1

	./configure --prefix=/usr/local/postgresql/ --with-blocksize=8  --with-perl --with-python --with-tcl --with-openssl --with-pam --with-ldap --with-libxml --with-libxslt --enable-thread-safety --enable-dtrace --with-uuid=ossp 

	make 

	make install

	#添加uuid的拓展
	cd contrib/uuid-ossp/
	make
	make install

	#初始化数据库目录
	initdb -D  /usr/local/postgresql/data -U postgres -E utf-8 --locale=C

	#启动
	pg_ctl start


####添加开机启动

	#返回root账户
	exit
	cp /usr/local/src/postgresql-10beta1/contrib/start-scripts/linux /etc/init.d/postgresql
	#添加执行权限
	chmod +x /etc/init.d/postgresql
	
	#修改起动脚本中路径
	vim /etc/init.d/postgresql
	# Installation prefix
	prefix=/usr/local/postgresql/
	
	# Data directory
	PGDATA="/usr/local/postgresql/data"
	
	# Who to run the postmaster as, usually "postgres".  (NOT "root")
	PGUSER=postgres

	# 添加开机启动
	chkconfig --add postgresql
	chkconfig  postgresql on

#### 命令

	#启动
	pg_ctl start -D $PGDATA -l pgsql.log
	
	#查看状态
	pg_ctl status

	#关闭
	pg_ctl stop -D $PGDATA

	#查看所有数据库
	 psql -l

	#登录数据库
	psql [databasename] #默认数据库postgres

	#新建数据库
	CREATE DATABASE exampledb OWNER dbuser;

	#新建用户	
	CREATE USER dbuser WITH PASSWORD 'password';

	#数据库授权
	GRANT ALL PRIVILEGES ON DATABASE exampledb to dbuser;

	#退出
	\q


#### 设置远程登录
	
	#修改postgresql.conf （在postgres/data 目录下）
	listen_addresses = '*'

	#修改pg_hba.conf	
	host  all  all 0.0.0.0/0 md5

	#重启postgresql 服务
	sudo service postgresql restart	


参考：

- [https://www.hyahm.com/article/147.html](https://www.hyahm.com/article/147.html)
- [https://my.oschina.net/tashi/blog/189351](https://my.oschina.net/tashi/blog/189351)
- [http://www.ruanyifeng.com/blog/2013/12/getting_started_with_postgresql.html](http://www.ruanyifeng.com/blog/2013/12/getting_started_with_postgresql.html)
- [http://lazybios.com/2016/11/how-to-make-postgreSQL-can-be-accessed-from-remote-client/](http://lazybios.com/2016/11/how-to-make-postgreSQL-can-be-accessed-from-remote-client/)

