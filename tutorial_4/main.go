package main

import "fmt"

func main() {
	var ages [3]int = [3]int{21, 25}
	ages[1] = 24
	fmt.Println(ages, len(ages))

	names := [4]string{"Lupin", "Diop", "Pellegrini", "Claire"}
	fmt.Println(names, len(names))

	// Slices - It's an Array under the hood, but you don't have to specify the length and it's possible to append elements to it.

	scores := []float32{25.00, 12.5}
	scores = append(scores, 13)
	fmt.Println(scores, len(scores))

	// Slice Range

	rangeOne := names[1:2]  // Create a new slice with values from the names array between index 1 and 4
	rangeTwo := ages[1:]    // [24 0], 0 is the default value of non value is specified.
	rangeThree := names[:3] // [24 0], 0 is the default value of non value is specified.

	rangeOne = append(rangeOne, "Lénoard")

	fmt.Println(rangeOne, rangeTwo, rangeThree)

}
