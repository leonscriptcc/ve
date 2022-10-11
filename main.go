package main

import (
	"github.com/leonscriptcc/ve/gconfig"
	"github.com/leonscriptcc/ve/ve"
	"log"
)

func init() {
	// 读取配置文件
	err := gconfig.Load()
	if err != nil {
		log.Println("config.yml is invalid:", err.Error())
	}
}

func main() {
	//TODO 循环初始化采集点
	//TODO 启动http监听
	ve.AnalyzeDataSrc("./data.txt")
}
