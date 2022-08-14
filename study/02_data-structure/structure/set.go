package structure

import (
	"sync"
)

type Set struct {
	data map[int]struct{}
	cap int
	len  int
	sync.RWMutex
}

func NewSet(cap int) *Set {
	tmp := make(map[int]struct{}, cap)
	return &Set{cap: cap, data: tmp}
}

func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()

	s.data[item] = struct{}{}
	s.len = len(s.data)
}

func (s *Set) Remove(item int) {
	s.Lock()
	defer s.Unlock()

	_, ok := s.data[item]
	if !ok {
		return
	} else {
		delete(s.data, item)
	}
}

func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()

	_, ok := s.data[item]
	return ok
}

func (s *Set) Length() (v int) {
	v = s.len
	return
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()

	s.data = make(map[int]struct{}, s.cap)
	s.len = 0
}

func (s *Set) IsEmpty() bool {
	return s.len == 0
}

func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()

	l := make([]int, 0, s.len)
	for v := range s.data {
		l = append(l, v)
	}

	return l
}

