package main

import (
	"fmt"
	"os"

	"github.com/fxtlabs/date"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

func (b *bill) format() string {
	fs := fmt.Sprintf("Breakdown for Bill %v: \n", b.name)
	var total float64 = b.tip

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip: ", b.tip)

	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "total: ", total)

	return fs
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) save() {
	data := []byte(b.format())
	currentTime := date.Today().Local().UnixMilli()
	fileName := fmt.Sprintf("bills/%v-%v.txt", b.name, currentTime)

	err := os.WriteFile(fileName, data, 7777)
	if err != nil {
		panic(fmt.Sprintf("Panic: %v", err))
	}

	fmt.Println("Bill was saved.")
}
