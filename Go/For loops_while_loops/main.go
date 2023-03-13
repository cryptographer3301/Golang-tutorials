package main

import "fmt"

func main() {
	fmt.Println()

	x := 5
	for i := 0; i <= x; i++ {
		fmt.Printf("%d\n", i)
		if i == 3 {
			continue;
		}
		fmt.Printf("%d\n--------", i)
	}
}
