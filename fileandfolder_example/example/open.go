package fileandfolderexample

import (
	"fmt"
	"os"
)

func Open_example() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Print(err)
		return
	}

	fileSize := stat.Size()
	data := make([]byte, fileSize)
	count, err := file.Read(data)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("read %d bytes from %s\n", count, stat.Name())
	fmt.Println(string(data))
}
