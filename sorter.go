package util

type Sorter struct {
	LenFunc  func() int
	SwapFunc func(i, j int)
	LessFunc func(i, j int) bool
}

func (s Sorter) Len() int {
	return s.LenFunc()
}

func (s Sorter) Swap(i, j int) {
	s.SwapFunc(i, j)
}

func (s Sorter) Less(i, j int) bool {
	return s.LessFunc(i, j)
}
