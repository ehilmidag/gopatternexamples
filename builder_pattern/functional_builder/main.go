package main

import "fmt"

type Person struct {
	name     string
	position string
}
type personMod func(*Person)
type PersonBuilder struct {
	actions []personMod
}

// bu implementasyondaki amaç mevcut builderi bozmadan yeni builder funclar implemente edebilmektir personMod sayesinde
func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}
func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}
func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}
func main() {
	b := PersonBuilder{}
	p := b.Called("hilmi").WorksAsA("dev").Build()
	fmt.Println(*p)
}