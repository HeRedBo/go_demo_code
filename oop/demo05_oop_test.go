package main

import "fmt"

type Worker interface {
	Work(hours int) (product string)
	Rest()
}
/*
*Coder实现的Worker接口，coder := Coder{},coderPtr := &coder,var worker Worker，var workders []Worker
=>worker只能向*Coder断言
=>只有coderPtr才能给worker赋值，只有coderPtr才能加入workders
*/
type Coder struct {}
func (c *Coder) Work(hours int) (product string) {
	fmt.Println("劳资在工作")
	return "BUG"
}

func (c *Coder) Rest() {
	fmt.Println("劳资在休息")
}

func main() {
	coder := Coder{}
	coderPtr := &coder

	//只有指针实现了接口，所以只有指针对象才是接口实例
	var worker Worker
	worker = coder
	worker = CoderPtr
	fmt.Println(worker)

	// 只有指针实现了接口，所以只有指针对象才是接口实例
	workers := make([]Worker, 0)
	workers = append(workers, coder)
	workers = append(workers, CoderPtr)
	fmt.Println(workers)

	w1, ok := worker.(Coder)
	fmt.Println(w1, ok)
	w2, ok := worker.(*Coder)
	fmt.Println(w2, ok)

	switch worker.(type) {
	case Coder:
		fmt.Println("Coder")
	case *Coder:
		fmt.Println("*Coder")
	}

}
