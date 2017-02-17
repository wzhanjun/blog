## doskey

windows自定义快捷键

1. 新增bat

2. 添加命令内容

@doskey www=cd /d d:\jun\wwwroot\$1
@doskey ll=dir

3. 使用Win+R，输入regedit进入注册表，找到HKEY_LOCAL_MACHINE\Software\Microsoft\Command Processor，右键新建，字符串值，名为AutoRun，值为C:\cmd-alias.bat，保存退出。


[http://blog.csdn.net/qq285744011/article/details/51134905](http://blog.csdn.net/qq285744011/article/details/51134905)