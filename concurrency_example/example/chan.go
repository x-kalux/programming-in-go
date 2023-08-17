package concurrency_example

import (
	"fmt"
	"sync"
	"time"
)

func sum(c chan int, numbers ...int) {
	_sum := 0
	for _, num := range numbers {
		_sum += num
	}
	c <- _sum
}

func printer(c chan int) {
	number := <-c
	fmt.Println(number)
}

func Channel_example() {
	c := make(chan int)
	go printer(c) //goroutine จะรอจนกว่า c channel มีข้อมูลส่งเข้ามาจึงจะทำงาน
	go printer(c)
	go sum(c, 1, 5, 7, 9)
	go sum(c, 10, 30)

	time.Sleep(time.Second)
}

func generateInt(min, max int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := min; i < max; i++ {
		ch <- i
	}
}
func Channel_example2() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go generateInt(50, 100, ch, &wg)
	go generateInt(950, 1000, ch, &wg)
	go func() {
		wg.Wait()
		close(ch)
	}()
	sum := 0
	for i := range ch {
		// fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)
}
