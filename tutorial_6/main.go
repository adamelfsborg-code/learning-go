package main

import "fmt"

func main() {
	index := 0
	for index <= 5 {
		fmt.Println(index)
		index++
	}

	for index := 0; index <= 5; index++ {
		fmt.Println(index)
	}

	names := []string{"Lupin", "Lèonard", "Diop", "Claire", "Pellegrini"}
	lengthOfNames := len(names)

	for index := 0; index < lengthOfNames; index++ {
		fmt.Println(names[index])
	}

	for index, value := range names {
		fmt.Println(index, value)
	}

	for index, value := range names {
		fmt.Println(value)
		if index == (lengthOfNames - 1) {
			names = append(names, "Youssef")
		}
	}

	fmt.Println(names)

}
