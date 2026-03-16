package main

import (
	"fmt"
	"sync"
	"time"
)

func worker_with_chan(done chan bool, id int) {
	fmt.Printf("Working...(%d)\n", id)
	time.Sleep(2 * time.Second)
	done <- true
	fmt.Printf("Done %d\n", id)
}

func worker(id int) {
	fmt.Printf("Working...(%d)\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Done %d\n", id)
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	var wg sync.WaitGroup
	wg.Go(func() { worker(1) })
	wg.Go(func() { worker(2) })
	wg.Wait()
	fmt.Println("All workers done")

	//-------------------------------

	done := make(chan bool, 1)
	go worker_with_chan(done, 3)
	<-done

	fmt.Println("ALL DONE")

	//-------------------------------
	//Channel directions
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
