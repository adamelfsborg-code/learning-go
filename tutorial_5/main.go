package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// strings package

	greeting := "Greetings sir!"

	doesContain := strings.Contains(greeting, "!sir")

	changedGreeting := strings.Replace(greeting, "sir", "mate", 1)

	countCharInGretting := strings.Count(greeting, "e")

	splittedString := strings.Split(greeting, " ")

	if doesContain {
		fmt.Println("Its a match")
	}

	fmt.Println(changedGreeting, greeting, countCharInGretting, splittedString)

	// sort package

	ages := []int{10, 12, 70, 23, 55, 36, 25, 12}

	names := []string{"Lupin", "Diop", "Lèonard", "Pellegrini", "Claire"}

	sort.Slice(ages, func(i, j int) bool {
		return ages[i] < ages[j]
	})
	indexOne := sort.SearchInts(ages, 40)

	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})
	indexTwo := sort.SearchStrings(names, "Pell")
	fmt.Println(ages, indexOne, names, indexTwo)
}
