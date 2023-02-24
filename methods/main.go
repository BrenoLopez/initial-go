package main

import "fmt"

type Person struct {
	name     string
	lastName string
	age      uint8
	status   bool
}
type Category struct {
	name string
	pai  *Category
}

func (c Category) HasParent() bool {
	return c.pai != nil
}
func (p Person) String() string {
	return fmt.Sprintf("Olá, meu nome é %s %s e eu tenho %d anos.", p.name, p.lastName, p.age)
}
func (c *Category) SetFather(father *Category) {
	c.pai = father
}
func main() {
	cat := Category{name: "Categoria 1"}
	if !cat.HasParent() {
		fmt.Println("Nao tem pai")
	}
	cat2 := Category{name: "Categoria 1", pai: &Category{name: "Categoria com pai"}}

	if cat2.HasParent() {
		fmt.Println("tem pai")
	}
	p := Person{name: "Breno", lastName: "Lopes", age: 25}
	fmt.Println(p)

	cat.SetFather(&Category{name: "pai"})
	if cat.HasParent() {
		fmt.Println("tem pai")
		fmt.Println(cat.pai.name)
	}
}
