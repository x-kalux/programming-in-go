package main

import (
	coursemanager "xkalux/panicexample/challenge"
)

func main() {
	var course = coursemanager.New_course("Compro")
	coursemanager.StartProgram(course)
}
