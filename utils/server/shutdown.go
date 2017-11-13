package server

import (
	"fmt"
)

//系统关闭释放资源
//必须阻塞同步关闭所有服务，资源回收后才能返回
func shutdown() bool {
	fmt.Println("server shutdown")
	return true
}
