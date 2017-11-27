package service

import (
	"github.com/dbv/dspider/utils/log"
	_ "github.com/dbv/dspider/modinit"
	"github.com/dbv/dspider/biz"
	//"time"
	"fmt"
)

func Start() {
	log.Debug("服务启动...")

	//time.Sleep(time.Second * 10000)
	s := biz.NewDsipder()
	//log.Debug(biz.Grules)
	if e := s.LoadRules(biz.Grules); e != nil {
		log.Debug("解析失败!", e.Error())
	}
	//定义每个项
	urlmap := make(map[string]int)
	for index := 0; index < len(s.BizRule.Root.Sites); index++ {
		urlmap[s.BizRule.Root.Sites[index]] = s.FuncGetPageList(s.BizRule.Root.Sites[index])
		//s.GetUserList(s.BizRule.Root.Sites[index])
	}
	//fmt.Print("urlmap:",urlmap)
	for k, v := range urlmap {
		log.Debug("开始处理url:", k)
		for index := 1; index <= v; index++ {
			listpage := fmt.Sprintf("%s&page=%d", k, index)
			log.Debug(listpage)
			s.GetUserList(listpage)
		}
	}
	//log.Debug(s.BizRule)
	//s.DoWork()
}

func Stop() {

	log.Debug("服务关闭...")
}
