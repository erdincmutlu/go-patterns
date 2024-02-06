package main

import "fmt"

type Employee struct {
	Name         string
	Position     string
	AnnualIncome int
}

const (
	Developer = 1
	Manager   = 2
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "developer", 60000}
	case Manager:
		return &Employee{"", "manager", 80000}
	default:
		panic("unsupported role")
	}
}

func main() {
	m := NewEmployee(Manager)
	m.Name = "Sam"
	fmt.Println(m)
}
