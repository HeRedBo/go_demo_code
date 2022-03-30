package main

import "fmt"

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

type Test struct {
	Max int
}

func (t *Test) Println() {
	fmt.Println(t.Max)
}

func deferExec(f func()) {
	f()
}

func main() {
	var t *Test
	defer deferExec(t.Println)

	//t = new(Test)
}
