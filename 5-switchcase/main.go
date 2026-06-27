package main

import (
	"fmt"
	"time"
)

func main() {
	// var x = 2
	// switch x { // no extrra print case problem
	// case 1:
	// 	{
	// 		fmt.Print("wertyu")
	// 		fmt.Print("yuwdfuj")
	// 	}
	// case 2:
	// 	{
	// 		fmt.Print("2, wertyu")
	// 		fmt.Print("2, yuwdfuj")
	// 	}
	// default:
	// 	fmt.Println("Hello")
	// }

	fmt.Println(time.Now().Weekday(), time.Now().Format(time.TimeOnly))
	fmt.Println(time.Now().Weekday(), time.Now().Format(time.DateOnly))
	switch time.Now().Weekday() {
	case time.Saturday:
		{
			fmt.Println("sdfghj")
		}
	}
}
