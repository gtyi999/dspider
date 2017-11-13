package utils

import (
	"net/http"
	"strings"
	"encoding/json"
)

//根据新版 HTTP 协议规范定义的 错误返回

type Nerror struct {
	Status           int         `json:"-"`
	ErrorType        string      `json:"error,omitempty"`
	ErrorDescription string      `json:"msg,omitempty"`
	ErrorCode        int         `json:"error_code,omitempty"`
	ErrorUri         string      `json:"error_uri,omitempty"`
	ErrorData        interface{} `json:"error_data,omitempty"`
}


//状态码	含义	说明
//200	OK	请求成功
//201	CREATED	创建成功
//202	ACCEPTED	更新成功
//400	BAD REQUEST	请求的地址不存在或者包含不支持的参数
//401	UNAUTHORIZED	未授权
//403	FORBIDDEN	被禁止访问
//404	NOT FOUND	请求的资源不存在
//500	INTERNAL SERVER ERROR	内部错误

var (
	Success = NewNerror(200).SetStatus(200).SetDescription("成功")
	ErrorNoRight = NewNerror(1000).SetStatus(403).SetDescription("没有权限")
	ErrorBadParam = NewNerror(1001).SetStatus(403).SetDescription("参数错误")

)

func NewNerror(errorCode int) *Nerror {
	return &Nerror{Status: 400, ErrorCode: errorCode}
}

func (this *Nerror) Error() string {
	resp := this.ErrorType
	if (this.ErrorDescription != "") {
		resp += ": " + this.ErrorDescription
	}
	return resp
}

func (this *Nerror) StatusCode() int {
	return this.Status
}

func (this *Nerror) SetDescription(description string) *Nerror {
	this.ErrorDescription = description
	return this
}

func (this *Nerror) SetErrorCode(code int) *Nerror {
	this.ErrorCode = code
	return this
}

func (this *Nerror) SetErrorDescription(errorDescription error) *Nerror {
	if errorDescription != nil {
		this.ErrorDescription = errorDescription.Error()
	}
	return this
}

func (this *Nerror) SetUri(uri string) *Nerror {
	this.ErrorUri = uri
	return this
}

func (this *Nerror) SetStatus(status int) *Nerror {
	this.Status = status
	return this
}

func (this *Nerror) SetData(data interface{}) *Nerror {
	this.ErrorData = data
	return this
}

func NewError(t string) *Nerror {
	return &Nerror{Status: 400, ErrorType: t}
}

func HttpStatusError(status int) *Nerror {
	return &Nerror{Status: status, ErrorType: strings.ToLower(strings.Replace(http.StatusText(status), " ", "_", -1))}
}

func JSONResponse(w http.ResponseWriter, data interface{}) {
	if err, ok := data.(*Nerror); ok {
		JSONResponseVerbose(w, err.Status, nil, err)
	} else {
		JSONResponseVerbose(w, 200, nil, data)
	}
}

func JSONResponseVerbose(w http.ResponseWriter, status int, header http.Header, data interface{}) {
	if header != nil {
		for k, v := range header {
			for _, vv := range v {
				w.Header().Set(k, vv)
			}
		}
	}
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "Fri, 01 Jan 1990 00:00:00 GMT")
	w.Header().Del("Content-Length")

	if bs, ok := data.([]byte); ok {
		w.WriteHeader(status)
		w.Write(bs)
		return
	}
	if d, err := json.Marshal(data); err != nil {
		panic("Error marshalling json: %v:" + err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(status)
		w.Write(d)
		return
	}
}
