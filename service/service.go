package service

import (
	"github.com/dbv/dspider/utils/log"
	_ "github.com/dbv/dspider/modinit"
	"time"
)

var flag = true

var Grules string = `
{
  "root:": {
    "name": "CSDN_BLOG",
    "sites": [
      "http://m.blog.csdn.net/Column/Column?Channel=mobile&Type=New",
      "http://m.blog.csdn.net/Column/Column?Channel=enterprise&Type=New",
      "http://m.blog.csdn.net/Column/Column?Channel=cloud&Type=New",
      "http://m.blog.csdn.net/Column/Column?Channel=www&Type=New",
      "http://m.blog.csdn.net/Column/Column?Channel=system&Type=New",
      "http://m.blog.csdn.net/Column/Column?Channel=database&Type=New",
      "http://m.blog.csdn.net/Column/Column?Channel=web&Type=New",
      "http://m.blog.csdn.net/Column/Column?Channel=code&Type=New",
      "http://m.blog.csdn.net/Column/Column?Channel=software&Type=New",
      "http://m.blog.csdn.net/Column/Column?Channel=other&Type=New"
    ]
  },
  "step": {
    "method": [
      "GetPageList,1,1,pagelist",
      "GetUserlist,1,1,userlist",
      "GetRelation,1,1,relationlist"
    ]
  }
}`


func Start() {
	log.Debug("服务启动...")
	for flag == true {
		time.Sleep(time.Second * 100)
		log.Debug("...")
	}
}

func Stop() {
	flag = false
	log.Debug("服务关闭...")
}
