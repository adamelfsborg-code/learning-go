package main

import "fmt"

func main() {

	// Strings

	var name string = "Adam"
	var lastName = "Elfsborg"
	var email string
	phoneNumber := "07302070432"

	fmt.Println(name, lastName, phoneNumber)

	email = "adam.elfsborg@hotmail.com"

	fmt.Println(email)

	// Ints

	var age int = 21
	var dateInMonth = 18
	var month int
	var alias rune = 55
	year := 2023

	fmt.Println(age, dateInMonth, year, alias)

	month = 9

	fmt.Println(month)

	// Bits § Memory

	var pinCode int16 = 5432
	var victories int8 = 5
	var date int32 = 20230918
	var dateAndTime int64 = 2023082211272015

	var notSure uint8 = 2 // Cant be negative, so the range is from 0 to 255.
	var aliasTwo byte = 55
	var hmm uintptr = 2222222222222222222 // Cant be negative.

	hmm = 2

	fmt.Println(pinCode, victories, date, dateAndTime, notSure, hmm, aliasTwo)

	// Floats

	var bill float32 = 5.0
	var largeNUm float64 = 30487348734437474387437843.248943834743747943
	score := 1.5 // type float64
	fmt.Println(bill, largeNUm, score)

}
