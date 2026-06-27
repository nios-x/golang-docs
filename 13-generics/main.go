package main

import "fmt" 

//any
//interface{}
//comparable

func printer[T interface{}](arr []T) {
	for _, v := range arr {
		fmt.Println(v)
	}
} 
type stack[T any] struct{
	arr []T
}
func main() {
	printer([]int{234, 543, 56, 46, 45, 65, 64})
	printer([]any{"defr", 76})
	st := stack[int]{
		arr:[]int{34565,53456,36,45,6},
	}
	fmt.Println(st)
}