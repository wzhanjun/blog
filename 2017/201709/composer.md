## composer 

#### composer安装

1. [http://www.phpcomposer.com/](http://www.phpcomposer.com/)

2. [https://getcomposer.org](https://getcomposer.org)

	全局安装

	1. 下载composer.phar 
	2. mv composer.phar /usr/local/bin/composer
	3. chmod +x composer 

#### composer 镜像

1. [https://pkg.phpcomposer.com/](https://pkg.phpcomposer.com/)

	1. 全局配置
	
		composer config -g repo.packagist composer https://packagist.phpcomposer.com
	
	2. 当前项目
		
		"repositories": {
		    "packagist": {
		        "type": "composer",
		        "url": "https://packagist.phpcomposer.com"
		    }
		}