package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var n int
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		go sumNumber(&n, i, &mutex)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("total number", n)
}

func sumNumber(n *int, i int, mutex *sync.Mutex) {
	mutex.Lock()
	*n += i
	mutex.Unlock()
}
