package concurrency_example

import (
	"fmt"
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
