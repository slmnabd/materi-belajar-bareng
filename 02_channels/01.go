package main

import "fmt"

func main() {

	message := make(chan string)

	go func() { message <- "hey!" }()

	msg := <-message
	close(message)

	fmt.Println(msg)
}
