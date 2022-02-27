package main

import (
	"fmt"
	"sync"
)

func main() {
	var n int
	var mutex sync.Mutex
	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		group.Add(1)
		go sumNumber(&n, i, &mutex, &group)
	}

	group.Wait()
	fmt.Println("total number", n)
}

func sumNumber(n *int, i int, mutex *sync.Mutex, group *sync.WaitGroup) {
	defer group.Done()
	mutex.Lock()
	*n += i
	mutex.Unlock()
}
