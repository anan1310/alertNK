package main

import (
	"alarm_collector/initialize"
)

func main() {
	//初始化基础配置
	initialize.InitBasic()
	//启动服务
	initialize.RunServer()
}
