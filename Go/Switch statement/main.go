package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"runtime"
)

func main() {
	fmt.Println("input text:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read line: %s-\n", scanner.Text())

	fmt.Println("switch statement in golang")

	reader := bufio.NewReader(os.Stdin)
	input, Error := reader.ReadString('\n')
	number, _ := strconv.ParseInt(input, 10, 64)

	switch number {
	case '5':
		fmt.Printf("Congrats")

	case '4':
		fmt.Printf("Soo close")

	case '2':
		fmt.Printf("Error")

	default:
		fmt.Printf("%d", Error)
	}

	fmt.Printf("%s", runtime.Version())
}
