package main

import "fmt"

func main() {

	var x [6]int = [6]int{1, 2, 3, 4, 5, 6}
	var slice []int = x[1:2] // First element is the starting index : 

	LAST_ELEMENT := slice[len(slice)-1] // +1

	fmt.Println("hello", x,slice, LAST_ELEMENT)
	fmt.Println(len(slice), slice[:cap(slice)])
}
