package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	position  string
	Salary    int
	ManagerId int
}

var dilbert Employee

func EmployeeByID(id int) Employee {
	return dilbert
}

func PtrEmployeeByID(id int) *Employee {
	return &dilbert
}

func main() {
	position := &dilbert.position
	*position = "Senior" + *position

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.position += "(proactive team player)"
	fmt.Println(1, *position)

	tmp := EmployeeByID(1)
	tmp.position = "ab"
	fmt.Println(EmployeeByID(1).position)
	fmt.Println(tmp.position)
	fmt.Println(&tmp == &dilbert)
	fmt.Println(dilbert.position)

	ptrEmployee := PtrEmployeeByID(1)
	ptrEmployee.position = "cd"
	fmt.Println(PtrEmployeeByID(1).position)

	PtrEmployeeByID(1).position = "ab"
	fmt.Println(PtrEmployeeByID(1).position)
}
