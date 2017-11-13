package utils

import (
	"net/url"
	"log"
	"time"
	"strconv"
	"bytes"
	"fmt"
	"math/rand"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"math"
	"encoding/binary"
	"unsafe"
)

var (
	DATA_JSON = "application/json;charset=utf-8"
	DATA_TEXT = "text/html;charset=UTF-8"
)

//todo: 向指定手机发送指定内容
func SendSms(tel string, msg string) (error) {
	log.Print("待发消息:", msg)
	var reqdata map[string]string
	reqdata = make(map[string]string)
	//必填参数。用户账号
	reqdata["name"] = "331234958"
	reqdata["pwd"] = "***"
	reqdata["content"] = msg
	reqdata["mobile"] = tel
	reqdata["stime"] = ""
	reqdata["sign"] = "wen94"
	reqdata["type"] = "pt"
	reqdata["extno"] = "0"
	var urlparms string
	iFlag := 0
	for k, v := range reqdata {
		if iFlag != 0 {
			urlparms = urlparms + "&"
			iFlag = 1
		}
		urlparms = urlparms + k + "="
		tempUri, pErr := url.Parse(v)
		if pErr != nil {
			return pErr
		}
		urlparms = urlparms + tempUri.EscapedPath()
		iFlag = 1
	}
	urlparms = "http://sms.1xinxi.cn/asmx/smsservice.aspx?" + urlparms
	GetData(urlparms)
	log.Print("url:", urlparms)
	return nil
}

//todo 生成指定长度的随机数(数字)
func GenRandom(ilen int) (num string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < ilen; i++ {
		i := r.Intn(9)
		num = num + strconv.Itoa(i)
	}
	return num
}

//todo 生成指定长度的随机数(字母)
func GenRandomCharset(ilen int) (charset string) {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < ilen; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//todo 根据时间获取显示的时间内容
func GetTimeContext(ctime time.Time) string {
	tnow := time.Now()
	dtime := tnow.Sub(ctime)
	//fmt.Println("dtime:",dtime.Seconds())
	if dtime.Seconds() < 0 {
		return "刚刚"
	}
	if dtime.Seconds() < 60 {
		return dtime.String() + " 秒前"
	}
	if dtime.Seconds() < 60*60 {
		return strconv.Itoa(int(dtime.Minutes())) + " 分钟前"
	}
	if dtime.Seconds() < 60*60*24 {
		return strconv.Itoa(int(dtime.Hours())) + " 小时前"
	}
	if dtime.Seconds() < 60*60*24*2 {
		return "昨天"
	}
	if dtime.Seconds() < 60*60*24*3 {
		return "前天"
	}
	if dtime.Seconds() < 60*60*24*7 {
		return strconv.Itoa(int(dtime.Hours()/24)) + " 天前"
	}
	return ctime.Format("2006-01-02 15:04:05")
}

//循环递归加载指定路径下的网页文件到缓存，是否递归添加
func LoadStaticTpl() {
	//AppTpl["home.index"] = pongo2.Must(pongo2.FromFile(AppCfg.TplFiles + "/home/index.html"))
	//AppTpl["home.search"] = pongo2.Must(pongo2.FromFile(AppCfg.TplFiles + "/home/search.html"))
	//AppTpl["user.regist"] = pongo2.Must(pongo2.FromFile(AppCfg.TplFiles + "/user/regist.html"))

	//AppTpl["home.index"] = pongo2.Must(pongo2.FromFile(AppCfg.TplFiles + "/home/index.html"))

	return
}

//todo 生成返回json数据信息
func GenRetMsg(retcode int, retmsg string) string {
	retmap := make(map[string]interface{})
	retmap["retcode"] = retcode
	retmap["retmsg"] = retmsg
	b, _ := json.Marshal(retmap)
	return string(b)
}

//todo post模拟提交数据
func PostData(uri string, dtype string, data []byte) (rep []byte, err error) {
	//把post表单发送给目标服务器
	body := bytes.NewBuffer([]byte(data))
	res, err := http.Post(uri, dtype, body)
	if err != nil {
		fmt.Println("发送错误:", err.Error())
		return nil, err
	}
	result, err := ioutil.ReadAll(res.Body)
	rep = result
	res.Body.Close()
	return rep, nil
}

func GetData(uri string) (rep []byte, err error) {
	r, e := http.Get(uri)
	if e != nil {
		return nil, e
	}
	result, err := ioutil.ReadAll(r.Body)
	rep = result
	r.Body.Close()
	return rep, nil
}

//todo float32 -> byte
func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}

//todo byte -> float32
func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}

//todo float64 -> byte
func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}

//todo byte -> float64
func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

//整形转换成字节
func IntToBytes(n int) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int(tmp)
}

func IsBigEndian() bool {
	const N int = int(unsafe.Sizeof(0))
	x := 0x1234
	p := unsafe.Pointer(&x)
	p2 := (*[N]byte)(p)
	if p2[0] == 0 {
		return true
	} else {
		return false
	}
}

func IsAllNumber(input string) (b bool) {
	for index := 0; index < len(input); index++ {
		if !(input[index] >= 0 && input[index] <= 9) {
			return false
		}
	}
	return true
}
