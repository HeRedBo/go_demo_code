package main

import (
	"log"
	"sync"
	"test/worker"
	"time"
)

var langs = []string{
	"Golang",
	"PHP",
	"Javascript",
	"Python",
	"Java",
	"C++",
}

type langPrinter struct{
	lang string
}

func (m *langPrinter) Task() {
	log.Println(m.lang)
	time.Sleep(time.Second)
}


func main() {
	p := worker.New(2)

	var wg sync.WaitGroup
	wg.Add(5 * len(langs))

	for i:= 0; i < 6 ; i ++ {
		for _, lang := range langs {
			lp := langPrinter{lang}
			go func() {
				p.Run(&lp)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	p.Shutdown()
}