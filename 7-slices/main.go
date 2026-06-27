package main

import "fmt"

func main() {
	var nums []int
	var nums2 = []int{}
	if nums == nil {
		fmt.Println("the nums is currently nil")
	}
	nums = append(nums, 344, 56456)
	fmt.Println(nums, nums2)

	var li = make([]int, 5) //, cap
	fmt.Println(li, len(li), cap(li))
	li = append(li, 34)
	fmt.Println(li)

}
