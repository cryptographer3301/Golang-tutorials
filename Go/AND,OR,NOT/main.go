package main

import "fmt"

func main() {
	fmt.Println("Logical operators in golang")

	var x = 5
	var y = 6

	fmt.Printf("%d", x / y)

	val := true || false && true // or || and && not !=
	var number rune = 2147483647 // int32 like - max range
	var num byte = 255 // MAX RANGE
	fmt.Printf("%t\n", val)
	fmt.Printf("%d\n %d", number, num)

	// Complex data type in golang: the first number is the real and the second is the imaginary number like 1 + 2i
	var z complex64 = 1 + 2i
	var n complex64 = 3 + 4i

	fmt.Printf("%#v %T\n", z, z)
	fmt.Println(z * n)
	fmt.Println(z / n)

	EXP := 32 + 13i
	fmt.Printf("%v", EXP)

	fmt.Printf("%v %T\n", real(z), real(z))
	fmt.Printf("%v %T\n", imag(z), imag(z))

	// STRINGS 
	str := "crypto" // starts from 1
	fmt.Printf("%d %q", len(str), str[1])
}
