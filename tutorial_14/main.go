package main

import "fmt"

func main() {

	myBill := newBill("Elgiganten")

	myBill.addItem("tv", 15)

	myBill.updateTip(15)

	fmt.Printf("myBill: %v\n", myBill.format())
}
