package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var cn int

func main() {
	cn = 0

	wg := new(sync.WaitGroup)

	m := new(sync.RWMutex)

	for i := 0; i < 11; i++ {
		wg.Add(3)
		go read(wg, m)
		go read(wg, m)
		go read(wg, m)
		go read(wg, m)
		go read(wg, m)
		go read(wg, m)
		go read(wg, m)
		go write(wg, m)
		go write(wg, m)
		go write(wg, m)
		time.Sleep(1 * time.Second)
	}

	wg.Wait()
}

func read(wg *sync.WaitGroup, m *sync.RWMutex) {
	m.RLock()
	defer m.RUnlock()

	cnToString := strconv.Itoa(cn)
	text := "-------カウント中: " + cnToString + " -------"
	// fmt.Printf("\r%s", text)
	// Go prayground用
	fmt.Printf("\x0c %s", text)
}

func write(wg *sync.WaitGroup, m *sync.RWMutex) {
	m.Lock()
	defer m.Unlock()

	cn++
	wg.Done()
}
