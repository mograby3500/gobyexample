package main

import (
	"fmt"
	"sync"
)

func ping(messages chan<- string, id int) {
	fmt.Printf("pre - ping, %d\n", id)
	messages <- fmt.Sprintf("Ping %d", id)
	fmt.Printf("post - ping, %d\n", id)

}

func read(messages <-chan string, id int) {
	fmt.Printf("pre - read, %d\n", id)
	msg := <-messages
	fmt.Printf("Reader %d: %s\n", id, msg)
	fmt.Printf("post - read, %d\n", id)
}

func main() {
	messages := make(chan string)

	var wg sync.WaitGroup
	wg.Go(func() {
		read(messages, 5)
	})
	wg.Go(func() {
		ping(messages, 1)
	})
	wg.Go(func() {
		ping(messages, 3)
	})
	wg.Go(func() {
		ping(messages, 4)
	})

	wg.Go(func() {
		read(messages, 2)
	})
	wg.Go(func() {
		read(messages, 6)
	})

	// Buffered channels
	buffered_messages := make(chan string, 2)

	buffered_messages <- "buffered"
	buffered_messages <- "channel"

	fmt.Println(<-buffered_messages)
	fmt.Println(<-buffered_messages)
	wg.Wait()
}
