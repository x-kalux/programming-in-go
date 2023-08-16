package fileandfolderexample

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	Id     string
	Name   string
	Age    int
	Points int
}

func Write_Json_example() {
	students := []student{
		{Id: "112-0", Age: 69, Name: "chonnan", Points: 7},
		{Id: "112-1", Age: 89, Name: "prawith", Points: 5},
		{Id: "112-2", Name: "itim", Points: 99},
	}

	jsonString, _ := json.Marshal(students)
	os.Mkdir("json_example", 0750)
	file, err := os.Create("json_example/students.json")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()
	file.Write(jsonString)
}

func Read_Json_example() {
	bytes, err := os.ReadFile("json_example/students.json")
	if err != nil {
		fmt.Print(err)
		return
	}
	students := []student{}
	json.Unmarshal(bytes, &students)

	fmt.Println(students)
	fmt.Println("---------")
	fmt.Println(students[2])
}
