package main

import "fmt"

func duplicateInArray(arr []int) []int {
	visited := make(map[int]bool, 0)
	slice := make([]int, 0)
	for i:=0; i<len(arr); i++{
		if visited[arr[i]] == true{
			slice = append(slice,arr[i])
		} else {
			visited[arr[i]] = true
		}
	}
	return slice
}

func main() {
	fmt.Println(duplicateInArray([]int{1, 4, 7, 2, 2,7}))
	fmt.Println(duplicateInArray([]int{1, 4, 7, 2, 3,3}))
}