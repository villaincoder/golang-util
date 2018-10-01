package util

import (
	"reflect"
	"sort"
	"testing"
)

func TestSorter(t *testing.T) {
	arr := []int{5, 3, 9, 4}
	sort.Sort(Sorter{
		LenFunc: func() int {
			return len(arr)
		},
		SwapFunc: func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		},
		LessFunc: func(i, j int) bool {
			return arr[i] < arr[j]
		},
	})
	if !reflect.DeepEqual(arr, []int{3, 4, 5, 9}) {
		t.Fatal("sort error")
	}
	t.Log("arr", arr)
}
