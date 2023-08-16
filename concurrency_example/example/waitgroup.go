package concurrency_example

import (
	"fmt"
	"sync"
	"time"
)

func print(txt string, sleep time.Duration, wg *sync.WaitGroup) {
	// defer wg.Done()
	fmt.Println(txt)
	time.Sleep(time.Microsecond * sleep)
}

func WaitGroup_example() {
	var wg sync.WaitGroup
	wg.Add(2)
	go print("Hello", 2, &wg)
	go print("Hi", 4, &wg)
	wg.Wait()
	fmt.Println("Done !!")
}
