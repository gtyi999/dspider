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
	"fmt"
	"net/http"
	"os"
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
		path:=""
		var err error
		if username!=""{
			path, err= os.Getwd()
			path += "/" + "DownLoadHtml"
			path+="/"+username
			//path += "\\" + username
			log.Debug("path:",path)
			isExist,_:=PathExists(path)
			if !isExist {
				err = os.Mkdir(path, os.ModePerm)
				if err != nil {
					log.Debug("err:", err)
				} else {
					log.Debug("Create Directory OK!")
				}
			}
		}
		//start add by guanty 下载用户的文章列表
		//如 :http://m.blog.csdn.net/blog/index?username=KingEasternSun
		userUrl := fmt.Sprintf("http://m.blog.csdn.net/blog/index?username=%v", username)
		DownLoadUserArticle(userUrl,path)
		modinit.RedisInstance.Set(username, time.Now().Format(time.RFC3339Nano), 0)

	})
	log.Debug("遍历结束")
	//fmt.Println("hrefs:",hrefs)
}

func  DownLoadUserArticle(url string,name string){
	log.Debug("DownLoadUserArticle,url:",url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Debug("DownLoadUserArticle,err:", err.Error())
	}

	doc.Find("ul[class=\"colu_author_c\"]").Eq(0).Find("li").Each(func(i int, content *goquery.Selection) {
		urlDetail,_ := content.Find("a").Attr("href")
		log.Debug("urlDetail:",urlDetail)

		htmlname:=urlDetail[strings.LastIndex(urlDetail, "/")+1:]+".html"
		log.Debug("htmlname:",htmlname)
		savePath:=""
		if strings.TrimSpace(urlDetail) != "" {
			urlDetail="http://m.blog.csdn.net"+urlDetail
			savePath=name
			savePath+="/"+htmlname
			log.Debug("savePath:",savePath)
			GetDetailArticle(urlDetail,savePath)
		}
		savePath=""
	})

}

func GetDetailArticle(url string,filename string){
	resp,err:=http.Get(url)
	if err!=nil{
		log.Debug("GetDetailArticle,err:", err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode==http.StatusOK{
		log.Debug("resp.StatusCode:", resp.StatusCode)
	}
	buf:=make([]byte,1024000)
	//createfile
	//f,err1:=os.OpenFile(filename,os.O_RDWR|os.O_CREATE|os.O_APPEND,os.ModePerm)
	f,err1:=os.OpenFile(filename,os.O_RDWR|os.O_CREATE,os.ModePerm)
	if err1!=nil{
		panic(err1)
		return
	}
	defer f.Close()
	for{
		n,_:=resp.Body.Read(buf)
		if 0==n{
			break
		}
		f.WriteString(string(buf[:n]))
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}