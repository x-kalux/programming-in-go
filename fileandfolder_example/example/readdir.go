package fileandfolderexample

import (
	"fmt"
	"os"
)

func Readdir_example() {
	files, err := os.ReadDir(".")
	// files, err := os.ReadDir("C:")
	if err != nil {
		fmt.Print(err)
		return
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
