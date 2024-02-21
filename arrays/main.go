package main

import "fmt"

func main() {

	var array [2]int
	array2 := [2]int{3, 5}

	var slice []int
	slice = []int{2, 3, 4, 5, 5}

	fmt.Printf("length: %d, capacity: %d, content: %d\n", len(array), cap(array), array)
	fmt.Printf("length: %d, capacity: %d, content: %d\n", len(array2), cap(array2), array2)
	fmt.Printf("length: %d, capacity: %d, content: %d\n", len(slice), cap(slice), slice)

	mySlice := make([]int, 5, 10)
	var mySlice2 []int
	mySlice2 = []int{2, 4}

	fmt.Printf("length: %d, capacity: %d, content: %d\n", len(mySlice), cap(mySlice), mySlice)
	fmt.Printf("length: %d, capacity: %d, content: %d\n", len(mySlice2), cap(mySlice2), mySlice2)

	var extendedSlice []int
	for i := 0; i < 100; i++ {
		extendedSlice = append(extendedSlice, i)
		fmt.Printf("length: %d, capacity: %d\n", len(extendedSlice), cap(extendedSlice))
		//fmt.Printf("length: %d, capacity: %d, content: %d\n", len(extendedSlice), cap(extendedSlice), extendedSlice)

	}

	var s1, s2, s3 []int
	s1 = []int{1, 2, 3, 4, 5, 6, 7, 8}
	s2 = s1
	s3 = make([]int, len(s1))
	copy(s3, s1)

	fmt.Printf("length: %d, capacity: %d, content: %d\n", len(s1), cap(s1), s1)
	fmt.Printf("length: %d, capacity: %d, content: %d\n", len(s2), cap(s2), s2)
	fmt.Printf("length: %d, capacity: %d, content: %d\n", len(s3), cap(s3), s3)
	s1[2] = 3
	s1[3] = 3
	s1[4] = 3
	s1[5] = 3
	fmt.Printf("length: %d, capacity: %d, content: %d\n", len(s1), cap(s1), s1)
	fmt.Printf("length: %d, capacity: %d, content: %d\n", len(s2), cap(s2), s2)
	fmt.Printf("length: %d, capacity: %d, content: %d\n", len(s3), cap(s3), s3)

	s4 := append(s3, 91)
	fmt.Printf("length: %d, capacity: %d, content: %d\n", len(s4), cap(s4), s4)
	s4[2] = 5
	fmt.Printf("length: %d, capacity: %d, content: %d\n", len(s3), cap(s3), s3)
	fmt.Printf("length: %d, capacity: %d, content: %d\n", len(s4), cap(s4), s4)

}
