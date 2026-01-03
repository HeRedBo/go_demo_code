package attrfomat

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/goutil/dump"
	"testing"
)

func TestAttrFormat(t *testing.T) {
	var jsonStr = `{
    "attrs": [
        {
            "detail": [
                "黑色",
                "白色",
                "红色"
            ],
            "value": "颜色"
        },
        {
            "detail": [
                "11",
                "111"
            ],
            "value": "尺寸"
        },
        {
            "detail": [
                "a3",
                "a4"
            ],
            "value": "大小"
        }
    ]
}`

	//jsonStr = `{"attrs":[{"detail":["11","30","40"],"value":"尺寸"}]}`

	// golang 字符串转json对象
	//var data any
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	var formatDetailList []FormatDetail
	jsonByte, _ := json.Marshal(data["attrs"])
	json.Unmarshal(jsonByte, &formatDetailList)
	var detailList = AttrFormt(formatDetailList)
	dump.P(detailList)
	//dump.P(formatDetailList)
}
