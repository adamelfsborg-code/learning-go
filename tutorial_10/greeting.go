package main

import "fmt"

var points = []int{20, 90, 31, 12, 5}

func greetHeros(h []string, f func(string)) {
	for _, v := range h {
		f(v)
	}
}

func sayHello(n string) {
	fmt.Printf("Hello, %v \n", n)
}

func execute() {
	greetHeros(heros, sayHello)
}
