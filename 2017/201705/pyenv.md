## Python

### python 多版本安装


##### 安装pyenv

	curl -L https://raw.githubusercontent.com/yyuu/pyenv-installer/master/bin/pyenv-installer | bash

安装完成后，根据提示将如下语句加入到 ~/.bashrc 中:

	export PYENV_ROOT="$HOME/.pyenv"
	export PATH="$PYENV_ROOT/bin:$PATH"
	eval "$(pyenv init -)"


参考:

[http://seisman.info/python-pyenv.html](http://seisman.info/python-pyenv.html)
