package fileandfolderexample

import (
	"fmt"
	"os"
)

func Create_example() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()

	file.WriteString("Hello\n")
	file.WriteString("i am test.txt file")
}
