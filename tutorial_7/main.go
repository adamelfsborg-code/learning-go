package main

import "fmt"

func main() {
	age := 45

	fmt.Println(age >= 50)
	fmt.Println(age <= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	if age < 45 {
		fmt.Println("Age is less then 45")
	} else if age > 45 {
		fmt.Println("Age is larger then 45")
	} else {
		fmt.Println("Age is equal to 45")
	}

	names := []string{"Lupin", "Diop", "Claire", "Lèonard", "Youssef", "Pellegrini"}

	for index, value := range names {
		if index == 1 {
			fmt.Printf("%v is second in list %v \n", value, names)
			continue
		}

		if index > 3 {
			break
		}
		fmt.Println(index, value)
	}

}
