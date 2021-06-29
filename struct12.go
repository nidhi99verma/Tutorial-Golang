package main

import "fmt"

type Student struct {
	rollno int
	name   string
	marks  int
}

func main() {
	student1 := Student{12, "Nidhi", 99}
	fmt.Println(student1)
	//or
	var student2 Student = Student{12, "Nidhi", 99}
	fmt.Println(student2)
	fmt.Println(student2.name)

	student3 := Student{rollno: 12, marks: 99}
	fmt.Println(student3)
	fmt.Println(student3.name)
}
