package main

import "fmt"

func main() {
	// array considerado um slice, pois tem tamanho dinâmico, diferente de um array comum que tem tamanho fixo
	names := []string{"Breno", "Thaynara", "Maria", "Edila"}
	fmt.Println(names)
	names = append(names, "Rafael")
	fmt.Println(names, len(names), cap(names))
	///////////////////////////////////////
	//array incializado de tamanho fixo de 20
	newNames := make([]string, 0, 20)
	newNames = append(newNames, "Breno")
	fmt.Println(newNames, len(newNames), cap(newNames))

	//////////////////// hash table
	// não tem ordenação dos dados retornados
	ages := make(map[string]uint8)
	ages["Breno"] = 25
	ages["Thaynara"] = 25
	ages["Rafael"] = 31
	fmt.Println(ages["Breno"])

	value, exists := ages["Lucas"]
	fmt.Println(value, exists)
	////////////////////////////////////////

	type Person struct {
		name     string
		lastName string
		age      uint8
		status   bool
	}
	person := Person{
		name:     "Breno",
		lastName: "Lopes",
		age:      25,
		status:   true,
	}
	fmt.Println(person)

	type PersonWithFather struct {
		name     string
		lastName string
		age      uint8
		status   bool
		//nao é possivel referênciar uma struct dentro dela mesma diretamente,somente vai ponteiro
		father *PersonWithFather
	}
	personWithFather := PersonWithFather{name: "Breno", age: 25, lastName: "Lopes", status: true, father: &PersonWithFather{name: "Gilson", lastName: "Gomes", age: 45, status: true}}
	fmt.Println(personWithFather)
}
