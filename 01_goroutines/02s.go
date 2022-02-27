package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 100; i++ {
		go printValue(i)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Done")
}

func printValue(number int) {
	fmt.Println(number)
}
