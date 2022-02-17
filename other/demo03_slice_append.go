package main

import "fmt"

func main() {
	var res [][]int
	sum := []int{2}
	res = append(res, sum)
	fmt.Println(sum, res)
	res = append([][]int{sum}, res...)
	fmt.Println(sum, res)
}
