# python 一些问题


### No such file or directory

在文件头添加 #!/usr/bin/env python3 后直接运行文件

在直接运行python 文件时提示No such file or directory

###### 原因：
该脚本文件格式是 dos 格式 而非 unix 格式。 dos格式下，换行符是 CRLF 的问题，使得第一行变成了（CR 的 ascii 码是 015）

##### 解决办法：

在 vim 中，查看文件格式，并设置，保存即可

查看

:set ff
设置

:set ff=unix

##### 参考：
[https://segmentfault.com/q/1010000000680188](https://segmentfault.com/q/1010000000680188)

[https://www.liaohuqiu.net/cn/posts/python-script-in-bin-directory/](https://www.liaohuqiu.net/cn/posts/python-script-in-bin-directory/)


 
