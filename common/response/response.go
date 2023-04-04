package response

import (
	"blogs/common/errorx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
	Flag    bool        `json:"flag"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		switch e := err.(type) {
		case *errorx.CodeError:
			body.Code = e.Code
			body.Message = e.Message
			body.Data = e.Data
		default:
			body.Code = 1
			body.Message = err.Error()
		}
	} else {
		body.Code = 20000
		body.Message = "操作成功"
		body.Data = resp
		body.Flag = true
	}
	httpx.OkJson(w, body)
}
