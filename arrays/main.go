package main

import (
	"fmt"
	"sort"
)

func main() {

	slice := CreateSlice(0)
	slice.AddElement(3)
	slice.AddElement(5)
	slice.AddElement(1)
	slice.AddElement(0)
	slice.AddElement(14)
	slice.AddElement(3)
	slice.PrintSlice()
	tmp := sort.Reverse(slice)
	fmt.Println(tmp)
	slice.PrintSlice()
	//
	//arr := make(arrType, 5)
	////var arr [5]int
	//fmt.Println(arr)
	//arr[0] = 1
	//arr[1] = 2
	//arr[2] = 3
	//arr[3] = 4
	//arr[4] = 5
	//arr = AddElement(arr, 8)
	//fmt.Println("State 1")
	//fmt.Println(arr)
	//ModifyArray(arr)
	//
	//fmt.Println("Arr 2")
	//fmt.Println(arr)
	//fmt.Println("Arr 1")
	//fmt.Println(arr)

}

func ModifyArray(arr []int) {
	arr[2] = 300
}

type arrType []int

func AddElement(arr []int, elem int) []int {
	arr = append(arr, elem)
	return arr
}
