# crontab
golang crontab
## install
	cd src
	go get github.com/astaxie/beego
	go get gopkg.in/mgo.v2

## deploy crond
	1. 守护进程运行，可通过supervisor部署管理
	2. 接受cront的请求 到目标机器执行任务

## deploy cront
	1. 配置到crontab 1分钟执行一次 
	2. 解析任务时间，将可执行的任务push到crond
	
## deploy web
	1. cd src/web/
	2. bee pack   //此命令会把	conf static views 以及 编译好的二进制文件打包成 web.zip
	3. copy web.zip 到运行目标机器 解压 执行 ./web 即可运行