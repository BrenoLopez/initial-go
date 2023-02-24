package main

import "fmt"

type Document interface {
	Doc() string
}

type Person struct {
	Name   string
	Age    uint8
	Status bool
}

func (person Person) String() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", person.Name, person.Age)
}

func (fiscPerson FisicPerson) Doc() string {
	return fmt.Sprintf("my document is: %s", fiscPerson.DocumentNumber)
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

func (juridicPerson JuridicPerson) Doc() string {
	return fmt.Sprintf("my document is: %s", juridicPerson.cnpj)
}

func show(doc Document) {
	// para funcionar tem que declarar mais de uma variavel para assecao
	// if doc, ok := doc.(FisicPerson); ok {
	// 	fmt.Println(doc.Name)
	// } else if doc, ok := doc.(JuridicPerson); ok {
	// 	fmt.Println(doc.SocialReason)
	// }
	// assertion of type
	//convesao da linguagem ok ser booleano pra dizer a asseçao da condicional
	switch doc.(type) {
	case FisicPerson:
		fmt.Println(doc.(FisicPerson).Name)
	case JuridicPerson:
		fmt.Println(doc.(JuridicPerson).SocialReason)
	default:
		fmt.Println("Unknow type")
	}

	// retorna quem implementa doc
	fmt.Println(doc)
	// retorna o que foi implementado no metodo Doc
	fmt.Println(doc.Doc())
}
func main() {
	person := FisicPerson{
		Person{
			Name:   "Breno",
			Age:    25,
			Status: true,
		},
		"Breno",
		"Lopes",
		"1132491406",
	}
	show(person)
	fmt.Println()

	juridicPerson := JuridicPerson{
		Person{
			Name:   "Breno",
			Age:    25,
			Status: true,
		},
		"Logo ali Ltda",
		"24546846548",
	}
	show(juridicPerson)
}
