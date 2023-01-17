package api

import "github.com/fanke15/tool/pkg/basic"

type RespToJson struct {
	Field string `json:"field"`
	Type  string `json:"type"`
	Desc  string `json:"desc"`
}

func ResponseFormat(data interface{}, msg ...string) map[string]interface{} {
	var temp = make([]interface{}, basic.Zero)
	basic.UnMarshal(basic.Marshal(data), &temp)

	var r = map[string]interface{}{
		"data":  data,
		"msg":   basic.StrNull,
		"code":  basic.Zero,
		"count": len(temp),
	}
	if len(msg) > basic.Zero {
		r["msg"] = msg[basic.Zero]
	}
	return r
}
