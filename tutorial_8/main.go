package main

import (
	"fmt"
	"math"
)

func sayGretting(n string) {
	fmt.Printf("Hello, %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * (r * r)
}

func main() {
	fmt.Println("Tutorial 8")
	sayGretting("Adam")

	names := []string{"Lupin", "Clarie", "Salvator", "Pellegrini", "Youssef", "Diop"}

	cycleNames(names, sayGretting)

	area := circleArea(5.3)

	fmt.Printf("Circle is %0.3f \n", area)
}
