package main

import (
	"fmt"
	"time"
)

func receive(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			break
		}

		fmt.Println(name, i)
	}
	fmt.Println(name + " is done.")
}

func main() {
	ch := make(chan int, 20)

	go receive("goroutine 1", ch)
	go receive("goroutine 2", ch)

	for i := 0; i < 100; i++ {
		ch <- i
	}

	close(ch)
	time.Sleep(3 * time.Second)
}
