package service

import (
	"github.com/dbv/dspider/utils/log"
	_ "github.com/dbv/dspider/modinit"
	"github.com/dbv/dspider/biz"
	//"time"
	"fmt"
)

var flag = true



func Start() {
	log.Debug("服务启动...")
	for flag == true {
		//time.Sleep(time.Second * 10000)
		s := biz.NewDsipder()
		//log.Debug(biz.Grules)
		if e:=s.LoadRules(biz.Grules); e!=nil {
			log.Debug("解析失败!",e.Error())
		}
		//定义每个项
		urlmap := make(map[string]int)
		for index:=0; index<len(s.BizRule.Root.Sites); index++ {
			urlmap[s.BizRule.Root.Sites[index]] = s.FuncGetPageList(s.BizRule.Root.Sites[index])

			s.GetUserList(s.BizRule.Root.Sites[index])
			break
		}
		fmt.Print("urlmap:",urlmap)
		//for k,v := range urlmap {
		//	for index:=0; index<v; index++ {
		//		s.GetUserList(k,)
		//	}
		//
		//}




		flag = false
		//log.Debug(s.BizRule)
		//s.DoWork()
	}
}

func Stop() {
	flag = false
	log.Debug("服务关闭...")
}
