package main

import "fmt"

type Address struct {
	Street, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	person := Person{
		Name: "hilmi",
		Address: &Address{
			Street:  "kemal",
			City:    "tekirdag",
			Country: "turkiye",
		},
	}
	//burda hillheimi personda kopyaladığımızda aslında iki variableyi de aynı pointere baktırdık
	// böylece hillheime yaptığımız müdahale doğal olarak refere ettiği alanda personu da etkiledi
	//deep copy işte tam olarak bu anlarda lazım
	hillheim := person
	hillheim.Name = "Hill"
	hillheim.Address.Street = "mattah"

	fmt.Println(person, person.Address)
	fmt.Println(hillheim, hillheim.Address)
	// tüm elemanları tek tek kopyalıyoruz ki artık 2 farklı obje davranabilsinler
	//fakat bu çok zahmetli bir işlem bunu yapmanın daha pratik yollarını bulmalıyız
	svd := person
	svd.Address = &Address{
		Street:  person.Address.Street,
		City:    person.Address.City,
		Country: person.Address.Country,
	}
	svd.Address.Street = "bolu mahallesi"
	fmt.Println(svd, svd.Address)

}
