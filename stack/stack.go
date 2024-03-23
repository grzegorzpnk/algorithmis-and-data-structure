package stack

import "fmt"

var Stack []int

func Push(st []int, val int) {
	st = append(st, val)
}

func Print(st []int) {
	for _, v := range st {
		fmt.Println(v)
	}
}
