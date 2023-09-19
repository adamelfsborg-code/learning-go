package main

import "fmt"

func main() {
	fmt.Println("Hello There")

	menu := map[string]float64{
		"sallad":   44.99,
		"sandwich": 34.99,
	}

	fmt.Printf("menu: %v\n", menu)

	keys := make([]string, 0, len(menu))
	for k := range menu {
		keys = append(keys, k)
	}

	fmt.Printf("keys: %v\n", keys)

	for k, v := range menu {
		fmt.Println(k, v)
	}

	for k := range menu {
		delete(menu, k)
	}

	fmt.Printf("menu: %v\n", menu)

	phonebook := map[int]string{
		123456789: "Lupin",
		567890399: "Lèonard",
		233899038: "Diop",
		349834735: "Youssef",
		448397497: "Salvador",
	}

	fmt.Printf("phonebook: %v\n", phonebook)
	fmt.Printf("phonebook Lupin: %v\n", phonebook[123456789])

	phonebook[123456789] = "Assene"

	fmt.Printf("phonebook: %v\n", phonebook)

}
