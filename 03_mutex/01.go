package main

import (
	"fmt"
	"time"
)

func main() {
	var n int

	for i := 0; i < 100; i++ {
		go sum(&n, i)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("total number", n)
}

func sum(n *int, i int) {
	*n += i
}


