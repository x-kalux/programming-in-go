package fileandfolderexample

import (
	"fmt"
	"os"
)

func Readfile_example() {
	bytes, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(bytes))
}
