package coursemanager

import "fmt"

func handlePanic(c course) {
	msg := recover()
	switch msg {
	case "to much":
		fmt.Println("your number out of range")
		StartProgram(c)
	default:
		text := fmt.Sprintf(" '%v' Not Implemented", msg)
		panic(text)
	}
}
