package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getUserInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Printf("%v: ", prompt)
	resp, err := r.ReadString('\n')
	return strings.TrimSpace(resp), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, err := getUserInput("Enter name of bill", reader)

	if err != nil {
		log.Fatal("Error reading from terminal", err)
	}

	b := newBill(name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, err := getUserInput("Choose option (a - Add item, u - Update tip, s - Save Bill)", reader)

	if err != nil {
		log.Fatal("Error reading from terminal", err)
		os.Exit(1)
	}

	switch opt {
	case "a":
		addItemToBill(&b)
		break
	case "u":
		updateTip(&b)
		break
	case "s":
		saveBill(&b)
		os.Exit(1)
	default:
		fmt.Println("Wrong input. Accepted Inputs are (a, u and s)")
		promptOptions(b)
	}
	promptOptions(b)
}

func addItemToBill(b *bill) {
	reader := bufio.NewReader(os.Stdin)
	itemName, err := getUserInput("Enter name of the item", reader)

	if err != nil {
		log.Fatal("Error reading from terminal", err)
	}

	itemPrice, err := getUserInput("Enter price of the item", reader)

	if err != nil {
		log.Fatal("Error reading from terminal", err)
	}

	price, err := strconv.ParseFloat(itemPrice, 64)

	if err != nil {
		log.Fatal("Error reading from terminal", err)
	}

	b.addItem(itemName, price)
}

func updateTip(b *bill) {
	reader := bufio.NewReader(os.Stdin)
	resp, err := getUserInput("Enter tip", reader)

	if err != nil {
		log.Fatal("Error reading from terminal", err)
	}

	tip, err := strconv.ParseFloat(resp, 64)

	if err != nil {
		log.Fatal("Error reading from terminal", err)
	}

	fmt.Print(tip, b.tip)
	b.updateTip(tip)
}

func saveBill(b *bill) {
	fmt.Println(b.format())
	b.save()
}

func main() {

	myBill := createBill()
	promptOptions(myBill)

	fmt.Println(myBill)

}
