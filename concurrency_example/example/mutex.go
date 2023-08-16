package concurrency_example

import (
	"fmt"
	"sync"
	"time"
)

func increment(data *int, rwmutex *sync.RWMutex, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	defer rwmutex.Unlock()
	rwmutex.Lock()
	*data++
	fmt.Println(time.Since(start), " Increment to: ", *data)
}

func read(data *int, rwmutex *sync.RWMutex, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	defer rwmutex.RUnlock()
	rwmutex.RLock()
	fmt.Println(time.Since(start), " Data: ", *data)
}

func Mutex_example() {
	var rwmutex sync.RWMutex
	var wg sync.WaitGroup
	data := 10
	wg.Add(10)
	for i := 0; i < 5; i++ {
		go increment(&data, &rwmutex, &wg)
	}
	for i := 0; i < 5; i++ {
		go read(&data, &rwmutex, &wg)
	}
	wg.Wait()
	fmt.Println("Done !!")
}
