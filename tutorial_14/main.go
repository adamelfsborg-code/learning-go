package main

import "fmt"

func main() {

	myBill := newBill("Elgiganten")

	billFs := myBill.Format()

	fmt.Printf("myBill: %v\n", billFs)
}
