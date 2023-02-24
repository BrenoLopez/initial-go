package main

import (
	"log"
	"math/rand"
	"os"
)

var face, crown int

func lauchCoin(side string) (crown int, face int) {
	println("side: ", side)
	switch side {
	case "face":
		face++
	case "crown":
		crown++
	default:
		println("caiu em pé")
	}
	return
}

func main() {
	a, b := 10, 13
	if a > b {
		println("a é maior do que b")
	} else if a < b {
		println("a é menor do que b")
	} else {
		println("a é igual a b")
	}
	// forma 1
	file, err := os.Open("hello.txt")
	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)
	// forma 2 onde err esta disponivel somente no escopo do if
	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}
	println(string(data))

	var sideLaunch string

	if rand.Intn(2) == 0 {
		sideLaunch = "face"
	} else {
		sideLaunch = "crown"
	}

	println(lauchCoin(sideLaunch))
}
