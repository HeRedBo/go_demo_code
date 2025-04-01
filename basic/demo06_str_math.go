package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	licenseKeyFormatting("2-5g-3-J",2)
}

func licenseKeyFormatting(s string, k int) string {
	s = strings.ToUpper(s)
	arr := strings.Split(s, "-")
	str2 := strings.Join(arr,"")
	str_len := strings.Count(str2,"") -1
	fmt.Println(str_len/k)
	step_num :=  math.Ceil(float64(str_len)/float64(k))
	step_count := int(step_num)
	end_start := str_len
	start_start := str_len
	slice := make([]string, 0)
	for i := 1 ;i <= step_count; i ++ {
		start_start= end_start - k
		if start_start < 0 {
			start_start = 0
		}
		fmt.Printf("[%d:%d]",start_start,end_start)

		str := str2[start_start:end_start]
		slice = append(slice,str)
		//fmt.Println(str2[start_start:end_start])
		end_start = start_start
	}
	fmt.Println(len(slice))


	fmt.Println(slice)

	fmt.Println(str2) // 5
	fmt.Println(str2[0:1]) // 0
	fmt.Println(str2[1:3]) // 1 5 - 2 x 2
	fmt.Println(str2[3:5]) // 2 5 - 2 x 1
	fmt.Println(arr)
	return ""
}
