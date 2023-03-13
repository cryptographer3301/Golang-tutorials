package main

import "fmt"

func main() {
	fmt.Println("Conditions and boolean expressions")
	x := 5
	y := 6.2
	val := float64(x) < y // explicit x to float64
	fmt.Printf("%t", val)
}