package main

import (
	"log"
	"runtime"
	"github.com/dbv/dspider/utils/server"
	"github.com/dbv/dspider/service"
)

func main() {
	//服务CPU初始化
	runtime.GOMAXPROCS(runtime.NumCPU())
	cStop := make(chan bool, 1)
	go server.GracefullyStopSever(func() {
		log.Print("服务优雅停止")
		service.Stop()
		//业务停止处理函数
		cStop <- true
	})
	//服务启动处理函数
	service.Start()
	<-cStop
}
