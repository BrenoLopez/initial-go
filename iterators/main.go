package main

import (
	"fmt"
	"iter"
)

func rootIterator(yield func(item int) bool) {
	for i := range 3 {
		if !yield(i) {
			return
		}
	}
}

func Keys(m map[string]string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for k := range m {
			if !yield(k) {
				return
			}
		}
	}
}

func Entries(m map[string]string) iter.Seq2[string, string] {
	return func(yield func(string, string) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

func main() {

	rootIterator(func(item int) bool {
		fmt.Println(item)
		if item >= 3 {
			return false
		}

		return true
	})

	for item := range rootIterator {
		fmt.Println(item)
	}

	x := make(map[string]string)
	x["x"] = "string"

	for key, value := range Entries(x) {
		fmt.Println(key + " " + value)
	}

	Entries(x)(func(key, value string) bool {
		fmt.Println(key + " " + value)
		return true
	})
}
