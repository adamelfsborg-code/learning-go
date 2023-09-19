package main

import (
	"fmt"
	"strings"
)

func updateName(n string) string {
	return strings.ToUpper(n)
}

func updateMenu(y map[string]int) {
	y["coffee"] = 52
}

func main() {
	// Group A types -> strings, ints, floats, bools, arrays, structs (Non pointer types)
	name := "Adam"

	name = updateName(name)

	fmt.Printf("name: %v\n", name)

	// Group B types -> slices, functions, maps (Pointer types)

	menu := map[string]int{
		"pizza": 39,
	}

	updateMenu(menu)

	fmt.Printf("menu: %v\n", menu)

}
