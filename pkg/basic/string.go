package basic

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// AnyToStr 将任意类型转换为string类型
func AnyToStr(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch vt := value.(type) {
	case float64:
		key = strconv.FormatFloat(vt, 'f', NegativeOne, SixtyFour)
	case float32:
		key = strconv.FormatFloat(float64(vt), 'f', NegativeOne, SixtyFour)
	case int:
		key = strconv.Itoa(vt)
	case uint:
		key = strconv.Itoa(int(vt))
	case int8:
		key = strconv.Itoa(int(vt))
	case uint8:
		key = strconv.Itoa(int(vt))
	case int16:
		key = strconv.Itoa(int(vt))
	case uint16:
		key = strconv.Itoa(int(vt))
	case int32:
		key = strconv.Itoa(int(vt))
	case uint32:
		key = strconv.Itoa(int(vt))
	case int64:
		key = strconv.FormatInt(vt, Ten)
	case uint64:
		key = strconv.FormatUint(vt, Ten)
	case string:
		key = vt
	case []byte:
		key = string(vt)
	default:
		key = string(Marshal(value))
	}
	return key
}

// AnySliceToStr 字符串拼接
func AnySliceToStr(sep string, strs ...string) string {
	var (
		build strings.Builder
		total = len(strs)
	)
	for k, v := range strs {
		build.WriteString(v)
		if total-One > k {
			build.WriteString(sep)
		}
	}
	return build.String()
}

// RemoveStrDuplication 移除切片重复内容
func RemoveStrDuplication(data []string, remove string) []string {
	var val = make([]string, Zero)
	for _, v := range data {
		if v != remove {
			val = append(val, v)
		}
	}
	return val
}

// IsStrExist 判断元素是否存在
func IsStrExist(data []string, exist string) bool {
	return strings.Contains(strings.Join(data, StrUnderline), exist)
}

type jsonData struct {
	Field string `json:"field"`
	Type  string `json:"type"`
	Desc  string `json:"desc"`
}

// ParseStruct 将struct格式化
func ParseStruct(text string) []byte {
	var (
		data        = make([]jsonData, Zero)
		fieldStatus = false
		fieldKey    = map[string]string{}
	)

	textSlice := strings.Split(text, "\n")
	for _, v := range textSlice {
		temp := RemoveStrDuplication(strings.Split(v, StrSpace), StrNull)
		if len(temp) < Two {
			continue
		}
		if IsStrExist(temp, "struct") {
			fieldStatus = true
			continue
		}

		if fieldStatus {
			valTemp := jsonData{
				Field: strings.Trim(temp[Zero], "\t"),
				Type:  temp[One],
			}

			if key := strings.Index(v, "//"); key > Zero {
				valTemp.Desc = strings.Trim(v[key+Two:], StrSpace)
			}
			match, _ := regexp.MatchString("^[a-zA-Z0-9.]+$", valTemp.Type)
			if _, ok := fieldKey[valTemp.Field]; !ok && match {
				fieldKey[valTemp.Field] = valTemp.Field
				data = append(data, valTemp)
			}
		}
	}
	// 排序
	sort.Slice(data, func(i, j int) bool {
		return data[i].Field < data[j].Field
	})
	return Marshal(data)
}
