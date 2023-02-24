package main

import "fmt"

type Person struct {
	Name   string
	Age    uint8
	Status bool
}

func (person Person) String() string {
	return fmt.Sprintf("Óla, meu nome é %s e eu tenho %d anos.", person.Name, person.Age)
}

type FisicPerson struct {
	Person
	Name           string
	LastName       string
	DocumentNumber string
}

type JuridicPerson struct {
	Person
	SocialReason string
	cnpj         string
}

func (fisicPerson FisicPerson) String() string {
	return fmt.Sprintf("Óla, meu nome é %s e eu tenho %d anos.", fisicPerson.Name, fisicPerson.Person.Age)
}

func (juridicPerson JuridicPerson) String() string {
	return fmt.Sprintf("Óla, meu nome é %s e eu tenho %d anos.", juridicPerson.Name, juridicPerson.Age)
}
func main() {
	p := FisicPerson{
		Person{
			Name:   "Breno",
			Age:    25,
			Status: true,
		},
		"Nmae",
		"Lopes", "11324959092",
	}
	x := JuridicPerson{
		Person{
			Name:   "Thaynara",
			Age:    25,
			Status: true,
		},
		"razao social",
		"11344533212",
	}
	fmt.Println(p)
	fmt.Println(x)
}
