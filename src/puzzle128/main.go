package main

import "sync"

//实现set

// set类型为bool，有点骚
type inter interface{}
type Set struct {
	m map[inter]bool
	sync.RWMutex
}

func New() *Set {
	return &Set{
		m: map[inter]bool{},
	}
}
func (s *Set) Add(item inter) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func main() {
	
}
