package biz

import (
	"encoding/json"
	"github.com/dbv/dspider/model"
	"github.com/dbv/dspider/utils"
	"github.com/dbv/dspider/utils/log"
	"strings"
)

var BizFuncMap map[string]interface{}  = nil

func init() {
	if BizFuncMap == nil {
		//注册公共操作函数
		BizFuncMap = make(map[string]interface{})
		BizFuncMap["GetPageList"] = NewFuncGetPageList()
		BizFuncMap["GetUserlist"] = new
	}
}

type BizDspider struct {
	BizRule model.Rule
}

func NewDsipder() *BizDspider {
	return &BizDspider{}
}

func (this *BizDspider) LoadRules(rule string) error {
	return json.Unmarshal([]byte(rule), &this.BizRule)
}

func (this *BizDspider) DoWork() error {
	log.Info("start to scan:",this.BizRule.Root.Name)
	//获取根目录网站代码
	for index:=0; index<len(this.BizRule.Root.Sites); index++ {
		b,e := utils.GetData(this.BizRule.Root.Sites[index])
		if e!=nil {
			log.Error("get url failed:",this.BizRule.Root.Sites[index])
		}
		go DoStep(b,this.BizRule)
	}
	return nil
}

func DoStep(data []byte,rule model.Rule) error {
	for index:=0; index<len(rule.Step.Method); index++ {
		//解析step数据
		pararms := strings.Split(rule.Step.Method[index],",")
	}
	return nil
}
