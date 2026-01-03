package attrfomat

import (
	"strings"
)

type FormatDetail struct {
	Detail []string `json:"detail"`
	Value  string   `json:"value"`
}

type Detail struct {
	Data []string                       `json:"data"`
	Res  []map[string]map[string]string `json:"res"`
}

// AttrFormt 组合SKU规则算法
func AttrFormt(formatDetailList []FormatDetail) Detail {
	var (
		data []string
		res  []map[string]map[string]string
	)
	if len(formatDetailList) > 1 {
		for i := 0; i < len(formatDetailList)-1; i++ {
			if i == 0 {
				data = formatDetailList[i].Detail
			}
			var tmp []string
			for _, v := range data {
				for _, g := range formatDetailList[i+1].Detail {
					rep2 := ""
					if i == 0 {
						rep2 = formatDetailList[i].Value + "_" + v + "-" + formatDetailList[i+1].Value + "_" + g
					} else {
						rep2 = v + "-" + formatDetailList[i+1].Value + "_" + g
					}
					tmp = append(tmp, rep2)
					// 如果是 最后一次遍历到最后一次时候 开始组转新数据
					if i == len(formatDetailList)-2 {
						var (
							rep4    = make(map[string]map[string]string)
							reptemp = make(map[string]string)
						)
						for _, h := range strings.Split(rep2, "-") {
							rep3 := strings.Split(h, "_")
							if len(rep3) > 1 {
								reptemp[rep3[0]] = rep3[1]
							} else {
								reptemp[rep3[0]] = ""
							}
						}
						rep4["detail"] = reptemp
						res = append(res, rep4)
					}
				}
			}
			if len(tmp) > 0 {
				data = tmp
			}
		}
	} else {
		// 一个规则时候
		var dataArr []string
		for _, formatdetail := range formatDetailList {
			for _, str := range formatdetail.Detail {
				var map2 = make(map[string]map[string]string)
				dataArr = append(dataArr, formatdetail.Value+"_"+str)
				map1 := map[string]string{
					formatdetail.Value: str,
				}
				map2["detail"] = map1
				res = append(res, map2)
			}
		}
		s := strings.Join(dataArr, "-")
		data = append(data, s)
	}
	return Detail{
		Data: data,
		Res:  res,
	}
}
