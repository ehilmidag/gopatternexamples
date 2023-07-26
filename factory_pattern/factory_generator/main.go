package main

import "fmt"

type Employee struct {
	Name, Position string
	Income         int
}

// functional implementasyonu
// Avantajları bir fonksiyonu ötekinin içine aktarabiliyoruz
// Böylece kullanıcının son şeklini verebildiği objeler türetebiliyoruz
func NewEmployeeFactory(position string, income int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, income}
	}
}

// structural implementasyonu
// bunun avantajı ise factory değerlerimizi modifiye edebiliyoruz. garsona biraz zam yapalım
type EmployeeFactory struct {
	Position string
	Income   int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.Income}
}

func NewEmployeeFactory2(position string, income int) *EmployeeFactory {
	return &EmployeeFactory{position, income}
}

func main() {
	doctorFactory := NewEmployeeFactory("doctor", 70000)
	policeFactory := NewEmployeeFactory("police", 40000)
	doctor := doctorFactory("msd")
	police := policeFactory("idris")
	fmt.Println(doctor)
	fmt.Println(police)
	garsonFactory := NewEmployeeFactory2("garson", 30000)
	garsonFactory.Income = 35000
	garson := garsonFactory.Create("kirve")

	fmt.Println(garson)

}
