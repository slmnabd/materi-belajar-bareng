package main

import (
	"fmt"
	"time"
)


func main() {
	channel := make(chan string)
	defer close(channel)
	fmt.Println("Printer Ready")
	go Reader2(channel)
	go Printer2(channel)

	time.Sleep(3 * time.Second)
}

func Reader2(channel chan<- string) {
	fmt.Println("Reading data...")
	time.Sleep(1 * time.Second)
	channel <- "Bumi itu datar kawan"
}

func Printer2(channel <-chan string) {
	data := <-channel
	fmt.Println("result:", data)
	fmt.Println("Done")
}
