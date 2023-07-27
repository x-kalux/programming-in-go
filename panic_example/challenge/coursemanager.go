package coursemanager

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type student struct {
	id     string
	name   string
	points int
}

type course struct {
	students map[string]student
	name     string
}

func New_course(name string) course {
	var _course course
	_course.students = make(map[string]student)
	_course.name = name
	return _course
}
func (c *course) del_student(std_id string) (error, bool) {
	_, found := c.students[std_id]
	if !found {
		return errors.New("id: " + std_id + " not exists"), false
	} else {
		delete(c.students, std_id)
	}
	return nil, true
}

func (c *course) add_student(std student) (error, bool) {
	_, found := c.students[std.id]
	if found {
		return errors.New("id: " + std.id + " already exists"), false
	} else {
		c.students[std.id] = std
	}
	return nil, true
}

func (c *course) get_student_by_id(std_id string) (error, student) {
	_, found := c.students[std_id]
	if found == false {
		return errors.New("student(id:" + std_id + ") not found"), student{}
	}
	return nil, c.students[std_id]
}
func (c *course) update_student(std student) (error, student) {
	_, found := c.students[std.id]
	if found == false {
		return errors.New("student(id:" + std.id + ") not found"), student{}
	}
	c.students[std.id] = std
	return nil, c.students[std.id]
}
func (c *course) toString() string {
	var text = fmt.Sprintf("course name: %v\n", c.name)
	text += "----------------------------\n"
	if len(c.students) < 1 {
		text += "student enrolls : 0"
		return text
	}
	for _, std := range c.students {
		text += fmt.Sprintf("  %v %v %v\n", std.id, std.name, std.points)
	}
	text += "----------------------------\n"
	text += fmt.Sprintf("student enrolls : %v\n", len(c.students))
	return text
}

func menuCreateNewStudent(c course) error {
	fmt.Println("\n1. Create a new student")
	var std_id, std_name string
	var std_points int
	fmt.Print("Enter(separate by space)\nStdID StdName and StdEmail: ")
	fmt.Scan(&std_id, &std_name, &std_points)
	var std = student{std_id, std_name, std_points}
	err, _ := c.add_student(std)
	if err != nil {
		return err
	}
	return nil
}
func menuUpdateStudentPoints(c course) error {
	fmt.Println("\n2. Update student points")
	var std_id string
	var std_points int
	fmt.Print("Enter(separate by space)\nStdID Points: ")
	fmt.Scan(&std_id, &std_points)
	err, std := c.get_student_by_id(std_id)
	if err != nil {
		return err
	}
	std.points += std_points
	c.update_student(std)
	return nil
}
func menuDeleteStudent(c course) error {
	fmt.Println("\n3. Delete student ")
	var std_id string
	fmt.Print("Enter StdID you want to delete: ")
	fmt.Scan(&std_id)
	err, _ := c.del_student(std_id)
	if err != nil {
		return err
	}
	return nil
}

func mainMenu() string {
	var input string
	fmt.Println("\n##### Course Manager #####")
	fmt.Println("  enter 1 to create new student")
	fmt.Println("  enter 2 to update student points")
	fmt.Println("  enter 3 to delete student")
	fmt.Println("  enter 4 to display course detail")
	fmt.Println("  enter q to close program")
	fmt.Println("---------------------------")
	fmt.Print(" > : ")
	fmt.Scan(&input)
	return input
}

func StartProgram(c course) {
	defer handlePanic(c)
	var input = mainMenu()
	switch strings.ToLower(input) {
	case "1":
		err := menuCreateNewStudent(c)
		if err != nil {
			fmt.Println("error: ", err)
			panic("cannot add student")
		}
	case "2":
		err := menuUpdateStudentPoints(c)
		if err != nil {
			fmt.Println("error: ", err)
			panic("cannot update student")
		}
	case "3":
		err := menuDeleteStudent(c)
		if err != nil {
			fmt.Println("error: ", err)
			panic("cannot delete student")
		}
	case "4":
		report := c.toString()
		fmt.Println(report)
	case "q":
		fmt.Printf("\n See you soon ♥\n bye...\n\n")
		os.Exit(0)
	default:
		fmt.Printf("Invalid command '%v' \n\n", input)
	}
	StartProgram(c)
}
