package main

import "fmt"

func change(s ...int) {
	s = append(s, 3, 4)
	fmt.Println("change", s)
}

var x = []int{2: 2, 3, 0: 1}

func main() {
	fmt.Println(x)
}
