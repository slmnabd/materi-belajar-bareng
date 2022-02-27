package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 100; i++ {
		printNumber(i)
	}

	fmt.Println("Done")
}

func printNumber(number int) {
	fmt.Println(number)
}
