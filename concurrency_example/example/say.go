package concurrency_example

import (
	"fmt"
	"time"
)

func say(text string) {
	for i := 0; i < 3; i++ {
		fmt.Println(time.Now(), " : ", i, " : ", text)
		time.Sleep(time.Second)
	}
}

func Say_example() {
	go say("Hello")
	go say("Hi")
	var input string
	fmt.Scan(&input)
}
