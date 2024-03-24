package main

import (
	"fmt"
	"sort"
)

type SliceInt []int

func (s SliceInt) Len() int           { return len(s) }
func (s SliceInt) Less(i, j int) bool { return s[i] < s[j] }
func (s SliceInt) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func CreateSlice(size int, capacity ...int) SliceInt {

	if len(capacity) > 0 {
		slice := make(SliceInt, size, capacity[0])
		fmt.Printf("Created slices, len: %d, capacity `%d\n", len(slice), cap(slice))
		fmt.Println(slice)
		return slice
	} else {
		slice := make(SliceInt, size)
		fmt.Printf("Created slices, len: %d, capacity %d\n", len(slice), cap(slice))
		fmt.Println(slice)
		return slice
	}
}

func (s SliceInt) PrintSlice() {
	fmt.Println(s)
}

// AddElement value receiver
// need to be done like that, cause append returns new slice, so we need to assign it under the address
func (s *SliceInt) AddElement(elem int) {
	*s = append(*s, elem)
}

func (s *SliceInt) DeleteLastElement() {
	*s = (*s)[:len(*s)-1]
}

func (s SliceInt) ReverseSlice() {
	//s = sort.Reverse(s)
	//sort.Slice()
}

func (s *SliceInt) SortSlice() {
	sort.Ints(*s)
}

func (s *SliceInt) RandomInitialization() {

}

func (s *SliceInt) CheckEquality() {

}
