package main

import (
	"testing"

	"github.com/gookit/goutil/dump"
)

func TestFanOutFanIn(t *testing.T) {
	in := generate(2, 3, 4, 5, 6, 7, 8)
	workers := fanOut(in, 3)
	out := fanIn(workers...)
	for n := range out {
		dump.P(n)
		t.Log(n)
	}
}
