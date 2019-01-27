# ssh


## 1. 生成密钥对

	ssh-keygen -t rsa -C "your_email@example.com"
	passphrase(密钥验证)


## 密钥登录

	# cat id_rsa.pub >> .ssh/authorized_keys
	

	$ vim /etc/ssh/sshd_config

	#禁用root账户登录，非必要，但为了安全性，请配置
	PermitRootLogin no

	#是否让 sshd 去检查用户家目录或相关档案的权限数据，这是为了担心使用者将某些重要档案的权限设错，可能会导致一些问题所致。例如使用者的 ~/.ssh/ 权限设错时，某些特殊情况下会不许用户登入
	StrictModes no

	# 是否允许用户自行使用成对的密钥系统进行登入行为，仅针对 version 2。至于自制的公钥数据就放置于用户家目录下的 .ssh/authorized_keys 内
	RSAAuthentication yes
	PubkeyAuthentication yes
	AuthorizedKeysFile %h/.ssh/authorized_keys

	# 有了证书登录了，就禁用密码登录吧，安全要紧
	PasswordAuthentication no

## 多帐号密钥


	创建 ~/.ssh/config
	# Default github user(first@mail.com)
	Host github.com
	HostName github.com
	User git
	IdentityFile C:/Users/username/.ssh/id_rsa

	# second user(second@mail.com)
	Host github-second
	HostName github.com
	User git
	IdentityFile C:/Users/username/.ssh/id_rsa_second



### 参考

[https://www.jianshu.com/p/fab3252b3192](https://www.jianshu.com/p/fab3252b3192)