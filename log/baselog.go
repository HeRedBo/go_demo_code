package log

import (
	"github.com/gookit/goutil/dump"
	"log"
)

func SimpleLog() {
	log.SetPrefix("[main]")
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lmsgprefix | log.Ldate | log.Ltime)
	dump.P(log.LstdFlags | log.Lmicroseconds | log.Lmsgprefix | log.Ldate | log.Ltime)
	log.Println("日志")
	log.Println("panic日志")
	log.Fatalln("错误日志")

}
