package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan string)
	go CallFriend(channel)

	msg := <-channel
	close(channel)

	fmt.Println(msg)
}

func CallFriend(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "Woi Roni"
}


