package test

import (
	"fmt"
	"github.com/fanke15/tool/pkg/basic"
	"strings"
	"testing"
)

type JsonData struct {
	Field string `json:"field"`
	Type  string `json:"type"`
}

var (
	fieldTypes = []string{"uint", "", "", "", "", "", "", "", ""}
)

// 测试将struct转换为json
func TestStructToJson(t *testing.T) {
	var (
		val         = make([]JsonData, 0)
		fieldStatus = false
	)

	text := "type (\n\t// 通用model字段\n\tCommonModel struct {\n\t\tID        uint      `gorm:\"column:id;primary_key\" json:\"id\"`         // 自增Id主键\n\t\tCreateAt  time.Time `gorm:\"column:created_at\" json:\"created_at\"`     // 创建时间\n\t\tUpdatedAt time.Time `gorm:\"column:updated_at\" json:\"updated_at\"`     // 更新时间\n\t\tDeleted   int       `gorm:\"column:deleted;default 0\" json:\"deleted\"` // 是否删除 [默认0:未删除,1:删除]\n\t}\n)\n\n"

	data := strings.Split(text, "\n")
	for _, v := range data {
		temp := basic.RemoveStrDuplication(strings.Split(v, basic.StrSpace), basic.StrNull)
		if len(temp) < basic.Two {
			continue
		}
		if basic.IsStrExist(temp, "struct") {
			fieldStatus = true
			continue
		}

		if fieldStatus {
			valTemp := JsonData{
				Field: strings.Trim(temp[basic.Zero], "\t"),
				Type:  temp[basic.One],
			}
			val = append(val, valTemp)
		}
	}
	fmt.Println(string(basic.Marshal(val)))
}
