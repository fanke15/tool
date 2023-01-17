package handler

import (
	"github.com/fanke15/tool/api"
	"github.com/fanke15/tool/pkg/basic"
	"github.com/kataras/iris/v12"
)

type ConvertHandler struct {
}

func (c *ConvertHandler) ToJson(ctx iris.Context) {
	var (
		req  = api.ReqToJson{}
		resp = make([]api.RespToJson, basic.Zero)
	)
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(api.ResponseFormat(nil, err.Error()))
		return
	}

	basic.UnMarshal(basic.ParseStruct(req.StructText), &resp)

	ctx.JSON(api.ResponseFormat(resp))
}
