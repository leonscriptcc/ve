package main

import (
	"github.com/leonscriptcc/ve/gconfig"
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

}
