package main

import "fmt"

type Person interface {
	SayHello()
}

type geekperson struct {
	name string
	age  int
}

type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi its %s, I am  %d shit old\n", p.name, p.age)
}
func (p *geekperson) SayHello() {
	fmt.Printf("Hi its %s, completely a geek", p.name)
}

func NewPerson(name string, age int) Person {
	if age > 25 {
		return &geekperson{
			name: name,
			age:  age,
		}
	}
	return &person{name, age}
}

func main() {
	p := NewPerson("hilmi", 29) // bu şekilde aslında bir person oluşturduk ama modifikasyona kapamış olduk
	// p.age := 30 gibi bir definasyon yapamayız çünkü person structu encapsüle bir halde yaşar
	p.SayHello()
	g := NewPerson("hillheim", 55)
	g.SayHello()
}
