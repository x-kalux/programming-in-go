package main

import (
	coursemanager "github.com/x-kalux/programming-in-go/panic-example/challenge"
)

func main() {
	var course = coursemanager.New_course("Compro")
	coursemanager.StartProgram(course)
}
