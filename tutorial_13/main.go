package main

import "fmt"

func updateName(n *string) {
	*n = "Adam"
}

func main() {
	name := "Elfsborg"

	m := &name

	updateName(m)

	fmt.Println("Memory location of name: ", m)

	fmt.Println("The value of the  memory location is: ", *m)

	fmt.Printf("name: %v\n", name)

}
