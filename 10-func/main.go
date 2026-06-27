package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func getlangs() (string, string) {
	return "Hello", "friend"
}

func call(fn func(...any) (int, error)) {
	fn("Call")
}

// closures
func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func adder(arr ...int) int {
	x := 0
	for _, v := range arr {
		x += v

	}
	return x

}
func changenum(no *int) {
	*no = 75
}

func main() {
	// result := add(43743, 45356)
	// fmt.Println(getlangs())
	// lan1, lan2 := getlangs()
	// fmt.Println(lan1, lan2)
	// call(fmt.Println)

	// counter := Counter()
	// fmt.Println(counter())
	// fmt.Println(counter())
	// fmt.Println(counter())

	// arr := []int{345, 34, 56, 3, 6, 46, 4}
	// fmt.Println(adder(arr...))

	var num int = 0
	changenum(&num)
	fmt.Println(num)
}
