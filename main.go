package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//? helper func for user input
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill () bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("Create a new bill, name: ", reader)

	b := newBill(name)
	fmt.Println("created bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("choose option (a - add item, s - save bill, t - add tip): ", reader)

	//? test
	// fmt.Println(opt)

	switch opt {
	case "a":
		// fmt.Println("you chose a")
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err  != nil {
			fmt.Println("price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("item added - ", name, p, "$")
		// fmt.Println(name, price)
		promptOptions(b)

	case "t":
		// fmt.Println("you chose t")
		tip, _ := getInput("enter the tip amount($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip added - ", t, "$")
		promptOptions(b)
		// fmt.Println(tip)
	case "s":
		fmt.Println("you chose to save bill")
		b.save()
	default:
		fmt.Println("that was not a valid option...")
		promptOptions(b)
	}
}


func main() {
	mybill := createBill()
	promptOptions(mybill)



	// fmt.Println(mybill)
}
