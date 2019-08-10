package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("goroutinesの処理が終わるの待ってるよ")

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go one(wg)
		go two(wg)
	}
	wg.Wait()

	fmt.Println("goroutinesの処理が終わったよ")
}

func one(w *sync.WaitGroup) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("oneを実行してるよ")
	w.Done()
}

func two(w *sync.WaitGroup) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("twoを実行してるよ")
	w.Done()
}
