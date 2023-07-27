package main

import "fmt"

type Employee struct {
	Name, Position string
	Income         int
}

// bu yöntemde farklı tiplere göre prototip üretim yapıyoruz
const (
	Dev = iota
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Dev:
		return &Employee{"", "dev", 45000}
	case Manager:
		return &Employee{"", "manager", 60000}
	default:
		panic("unknown role")
	}
}

func main() {
	manager := NewEmployee(Manager)
	manager.Name = "Hilmi"
	fmt.Println(manager)
}
