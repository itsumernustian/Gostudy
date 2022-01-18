package main

import "fmt"

func main() {
	fmt.Println("Umer Farooq")
	var emp employee
	emp.phone_no = 69
	emp.email = "umerfarooq1174@gmail.com"
	emp.name = "Muhammad Umer Farooq"
	display(emp)
}

type employee struct {
	name     string
	email    string
	phone_no int64
}

func display(e employee) {
	fmt.Println(e.name)
	fmt.Println(e.email)
	fmt.Println(e.phone_no)
}
