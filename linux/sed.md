# sed 命令

### nginx 日志分割

	sed -n '/01\/Dec\/2018:00:00:00/, /26\/Jan\/2019:17:00:00/p' access.log > 2019.log