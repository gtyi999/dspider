package main

import (
	"runtime"
	"github.com/dbv/dspider/service"
)

func main() {
	//服务CPU初始化
	runtime.GOMAXPROCS(runtime.NumCPU())
	service.Start()
}
