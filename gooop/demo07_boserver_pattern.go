package gooop

import "sync"

// 观察者接口
type Observer interface {
	Update(message string)
}

// 主题
type Subject struct {
	observers []Observer
	mu        sync.RWMutex
}

func (s *Subject) Attach(o Observer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.observers = append(s.observers, o)
}

func (s *Subject) Notify(message string) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, observer := range s.observers {
		observer.Update(message)
	}
}
