package main

import (
	"fmt"
	// "os"
)

func main() {
	// fmt.Println("hello crypto")

	// os.Exit(2)

	// fmt.Println("end crypto !")

	fmt.Printf("Hello cryptographer: %6x", 3.453)

	// To store fmt function inside a variable
	var name string = fmt.Sprintf("Bonsoir\n")
	fmt.Printf(name)

	if(name == "Bonsoir\n") {
		fmt.Printf("%t ", true)

	} else {
		fmt.Printf("%t", false)
	}
}
