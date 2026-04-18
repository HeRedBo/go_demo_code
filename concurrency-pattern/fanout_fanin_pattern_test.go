package main

import (
	"testing"
)

func TestFanOutFanIn(t *testing.T) {
	in := generate(2, 3, 4, 5, 6, 7, 8)
	workers := fanOut(in, 3)
	out := fanIn(workers...)
	for n := range out {
		t.Log(n)
	}

}

func TestRunSearchPipeline(t *testing.T) {

	runSearchPipeline("Golang 并发编程")
}

func TestRobustSearch(t *testing.T) {
	robustSearch("Golang 错误处理")
}
