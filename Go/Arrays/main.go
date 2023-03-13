package main

import "fmt"

func main() {
	fmt.Println("go")
	/*
		An array is a numbered sequence of elements of a single type with a fixed length. In
	Go, they look like this:
	*/

	var a [5]int
	a[1] = 1
	fmt.Printf("%d", a)

	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Printf("%f", total / float64(len(x))) // casting len(x) of type int to float64 type .

	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	  }
	  for key, value := range m {
		  fmt.Println(key, value)
	  }
}
