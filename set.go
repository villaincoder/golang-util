package util

var empty = struct{}{}

type SetForeachFunc func(index int, value interface{}) bool

type Set struct {
	m map[interface{}]struct{}
}

func NewSet(items ...interface{}) *Set {
	set := &Set{}
	set.m = make(map[interface{}]struct{})
	set.Add(items...)
	return set
}

func (s *Set) Add(items ...interface{}) {
	for _, item := range items {
		s.m[item] = empty
	}
}

func (s *Set) Remove(item interface{}) {
	delete(s.m, item)
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) Clear() {
	s.m = make(map[interface{}]struct{})
}

func (s *Set) Foreach(f SetForeachFunc) {
	index := 0
	for k, _ := range s.m {
		if f(index, k) {
			break
		}
		index++
	}
}
