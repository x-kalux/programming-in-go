package fileandfolderexample

import (
	"fmt"
	"os"
)

func Mkdir_example() {
	err := os.Mkdir("example/testdir", 0750)
	if err != nil {
		fmt.Print(err)
		return
	}
	err = os.WriteFile("example/testdir/testfile.txt",
		[]byte("Hello World"), 0660)
}
