package biz

import (
	"encoding/json"
	"github.com/dbv/dspider/model"
	"github.com/dbv/dspider/utils/log"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"strconv"
	"github.com/dbv/dspider/modinit"
	"time"
)

var Grules string = `
{
  "root": {
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

type BizDspider struct {
	BizRule model.Rule
}

func NewDsipder() *BizDspider {
	return &BizDspider{}
}

func (this *BizDspider) LoadRules(rule string) error {
	return json.Unmarshal([]byte(rule), &this.BizRule)
}

func (this *BizDspider) FuncGetPageList(url string) (int) {
	//log.Debug("start to parse url:", url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Debug("初始化文档失败", err.Error())
	}
	pagestr := doc.Find(".page_nav span").Text()
	if len(pagestr) < 4 {
		return -1
	}
	//fmt.Println("pagestr:", pagestr)
	nstart := strings.LastIndex(pagestr, "共")
	nend := strings.LastIndex(pagestr, "页")
	pagecount, err := strconv.Atoi(pagestr[nstart+3:nend])
	if err != nil {
		return -1
	}
	//log.Debug("总共:",pagecount)
	return pagecount
}

func (this *BizDspider) GetUserList(url string) () {
	//存入数据库
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Debug("初始化文档失败", err.Error())
	}
	doc.Find(".colu_b").Each(func(i int, content *goquery.Selection) {
		username := content.Find(".my_name").Text()
		log.Debug("username:", username)
		modinit.RedisInstance.Set(username, time.Now().Format(time.RFC3339Nano), 0)

	})
	log.Debug("遍历结束")
	//fmt.Println("hrefs:",hrefs)
}
