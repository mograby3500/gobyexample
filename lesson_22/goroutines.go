package main

import (
	"fmt"
	"sync"
	"time"
)

func f(from string) {
	for i := range 5 {
		fmt.Println(from, ":", i)
	}
}

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	go f("goroutine1")
	go f("goroutine2")
	go f("goroutine3")
	time.Sleep(time.Second)

	//WaitGroup

	var wg sync.WaitGroup

	for i := range 5 {
		wg.Go(func() {
			worker(i)
		})
	}
	wg.Wait()
}
