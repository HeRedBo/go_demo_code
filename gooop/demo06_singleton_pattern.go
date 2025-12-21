package gooop

import "sync"

type Singleton struct {
	data string
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{data: "initial data"}
	})
	return instance
}
