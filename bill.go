package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// ? create a new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// ? format the bill
func (b *bill) format() string {
	fs := "bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v...$%v \n", k+":", v)
		total += v
	}
	total += b.tip
	fs += fmt.Sprintf("%-25v...$%v \n", "tip: ", b.tip)

	//? total
	fs += fmt.Sprintf("%-25v...$%v \n", "total: ", total)
	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip //! this wont work as we are just changing the calues of the lcoal copy of the bill
}
//? for above func to work we have to involve pointers,we just have to add * beofre bill in func signature

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

//?save bill
func (b *bill) save () {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("bill was saved to file bills/"+b.name+".txt" )
} 