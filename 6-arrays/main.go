package main

import "fmt"

func main() {
	var nums [4]int
	nums[2] = 34
	fmt.Println(nums)
	fmt.Println(len(nums))
	nums2 := [3]int{2, 34, 45}
	fmt.Println(nums2)
	nums3 := [2][2]int{{2, 3}, {34, 45}}
	fmt.Println(nums3)

}
