package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {

	if strings.Trim(n, "") == "" {
		return "_", "_"
	}

	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	if len(initials) == 1 {
		return initials[0], "_"
	}

	return "_", "_"

}

func main() {
	firstChar, secondChar := getInitials("")

	fmt.Printf("%v.%v \n", firstChar, secondChar)
}
