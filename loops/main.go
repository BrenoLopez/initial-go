package main

import "fmt"

func main() {
	names := []string{"Breno", "Thaynara", "Maria", "Edila"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	// for of do go
	for _, name := range names {
		fmt.Println(name)
	}

	//equivalente ao while
	var i int
	for i < len(names) {
		fmt.Println(names[i])
		i++
	}
	// loop infinito equivalente a while(true)
	// for {
	// 	fmt.Println("print")
	// }
}
