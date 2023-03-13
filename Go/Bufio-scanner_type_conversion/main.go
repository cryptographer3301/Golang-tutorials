package main

import (
	"fmt"
)

func main() {
	// a := 1
	// fmt.Println("go")

	// for a <= 10 {
	// 	fmt.Printf("%d", a)
	// 	a++

	// 	if a % 2 == 0 {
	// 		fmt.Printf("%d", a)
	// 	}else {
	// 		continue
	// 	}
	// }

	number := 1
	for number <= 20 {
		if number % 3 == 0 {
			fmt.Printf("%s%d%s", "Feez", number, "\n" )
		} else if number % 5 == 0 {
			fmt.Printf("%s%d", "Buzz\n", number)
		}
		number++
	}
}
