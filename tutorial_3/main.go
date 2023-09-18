package main

import "fmt"

func main() {
	name := "Adam"
	age := 21
	fmt.Print("Hello,  ")
	fmt.Print("World! \n")
	fmt.Print("New Line! \n")

	// fmt.Printf (formatted strings) %_ = format specifier
	// %v = standrd.
	// %q = quote - Used for strings.
	// %T = type
	// %0.{decimals}f = float
	// %+d = show sign.
	fmt.Printf("My name is %v and I'm %v year's old! \n", name, age)
	fmt.Printf("My name is %q and I'm %q year's old! \n", name, age)
	fmt.Printf("Age is of type '%T' \n", age)
	fmt.Printf("Your highscore is %0.2f \n", 225.18)
	fmt.Printf("Your highscore is %b \n", 10)

	baseTwo := fmt.Sprintf("Base twop of int %v is %b", age, age)

	fmt.Println(baseTwo)

}
