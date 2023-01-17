package handler

import (
	"github.com/fanke15/tool/api"
	"github.com/fanke15/tool/pkg/basic"
	"github.com/kataras/iris/v12"
)

type DictHandler struct {
}

func (d *DictHandler) GetDickFieldList(ctx iris.Context) {
	var (
		data = "[{\"field\":\"Avax\",\"type\":\"decimal.Decimal\",\"desc\":\"amount to supply\"},{\"field\":\"AvaxAaveV3\",\"type\":\"decimal.Decimal\",\"desc\":\"\"},{\"field\":\"Btcb\",\"type\":\"decimal.Decimal\",\"desc\":\"\"},{\"field\":\"BtcbAaveV3\",\"type\":\"decimal.Decimal\",\"desc\":\"\"},{\"field\":\"Day\",\"type\":\"string\",\"desc\":\"用户id\"},{\"field\":\"Savax\",\"type\":\"decimal.Decimal\",\"desc\":\"amount to supply\"},{\"field\":\"SavaxAaveV3\",\"type\":\"decimal.Decimal\",\"desc\":\"\"},{\"field\":\"Wbtce\",\"type\":\"decimal.Decimal\",\"desc\":\"\"},{\"field\":\"WbtceAaveV3\",\"type\":\"decimal.Decimal\",\"desc\":\"\"}]"
		resp = make([]api.RespToJson, basic.Zero)
	)

	basic.UnMarshal([]byte(data), &resp)

	ctx.JSON(api.ResponseFormat(resp))
}
