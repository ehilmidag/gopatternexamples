package main

type Person struct {
	Name     string
	Age      int
	EyeCount int // göz genellikle 2 dir o yüzden customize etmeyiz çok factory function burda devreye girer
}

func NewPerson(name string, age int) Person {
	if age < 18 {
		panic("kardeşim büyü de gel") // böyle kontrolelri de factory function içinde kullanabiliriz
	}
	return Person{name, age, 2}
}

func main() {
	p := NewPerson("hilmi", 28)
	p.EyeCount = 3 // sadece mutasyona uğradığımda özel olarak manupile ederim
}
