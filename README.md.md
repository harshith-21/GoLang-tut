# **GO**

### TEMPLATE


```go
package main

import "fmt"

func main(){
	fmt.Println("Hello world")	
}
```

```go
// multiple packages
import (
	"fmt"
	"math"
	"strings"
)
```

### **VARIABLES**

```go
package main

import "fmt"

func main(){
	//? variables
	var nameOne string = "apple"

	fmt.Println(nameOne)
}
```

### **STRINGS**
```go
package main

import "fmt"

func main(){
	//? variables
	var nameOne string = "apple"
	var nameTwo = "ball"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameTwo = "cat"
	nameThree = "dog"

	fmt.Println(nameOne, nameTwo, nameThree)

	//? we can also do
	//? we cannot use this outside the function
	nameFour := "elephant"
	fmt.Println(nameFour)

}
```

### **NUMBERS**

```go
func main(){
	//? int

	var numOne int = 22
	var numTwo = 55
	numThree := 44

	fmt.Println(numOne, numTwo, numThree)

	//! int8 int16 int32 int64
	//! uint8 uint16 uint32 uint64

	var numFour int8 = 16
	var numFive uint16 = 12412 
	var numSix uint32 = 1241234645
	var numSeven uint32 = -4645

	//? float
	var scoreOne float32 = 123.34534
	var scoreTwo float64 = 12456456412423.235346457
}
```

### **STRINGS**

```go
	//* print
	fmt.Print("hello, ") //? does not send curser to new line
	fmt.Print("world") //? does not send curser to new line
	//! print on same line: hello, world

	//? for new line we write "\n" similar to cpp
	fmt.Print("\none\ntwo\nthree\n")

	//* println
	fmt.Println("one")
	fmt.Println("two")
	fmt.Println("three")

	name := "Harshith"
	age := 22
	
	fmt.Println("hello there, my name is ", name, "and age is ", age)

	//? in a better way
	//* formatted string
	fmt.Printf("my name is %v and my age is %v \n", name, age) 
	//? %v is default for all variables
	//? %_ is a format specifier: _ could be many diff letters acc to usage
	fmt.Printf("my name is %q and my age is %q \n", name, age) //? %q adds quotes to strings, wont work properly for others

	//? this prints the type of variable
	fmt.Printf("age is of type %T \n", age)

	//? for floats
	fmt.Printf("this is decimal: %f \n", 44.44)
	fmt.Printf("this is decimal: %0.3f \n", 44.44)

	//? Sprintf ( save formatted strings)
	
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("this is the saved string: ", str)
```
```
hello there, my name is  Harshith and age is  22
my name is Harshith and my age is 22
my name is "Harshith" and my age is '\x16'
age is of type int
this is decimal: 44.440000
this is decimal: 44.440
this is the saved string:  my age is 22 and my name is Harshith
```


### **ARRRAYS and SLICES**

#### **ARRAYS**
```go
	//? SLICES abnd arrays

	//? array
	var ages [3]int = [3]int{20, 21, 22} //? fixed length

	//? below works same as above
	// var agess  = [3]int{20, 21, 22} 

	fmt.Println(ages, len(ages))
	

	names := [4]string{"harshi", "hars", "haha", "hehe"}
	fmt.Println(names, len(names))

	//?eg 
	for i := 0; i < 3; i++ {
		fmt.Println("name: ", names[i], "age: ", ages[i])
	}
```

```
[20 21 22] 3
[harshi hars haha hehe] 4
name:  harshi age:  20
name:  hars age:  21
name:  haha age:  22
```

#### **SLICES**
```go
	//? SLICES use arrays under the hood (similar to vector in cpp)
	var scores = []int{1, 2, 3, 4}    //? no number in [] means slice

	// var scores1 = append(scores, 6)
	scores = append(scores, 5)
	scores[1] = 9

	fmt.Println(scores)
	fmt.Println(append(scores, 6))
	fmt.Println(scores)

	fmt.Println(" ")

	//? SLICE RANGES
	rangeOne := scores[1:3]
	fmt.Println(rangeOne)
	fmt.Println(scores[3:])
	fmt.Println(scores[0:])
	fmt.Println(scores[:3])
	fmt.Println("inclusive of first index, excl of second index: ",scores[1:3])
```
```
[1 9 3 4 5]
[1 9 3 4 5 6]
[1 9 3 4 5]

[9 3]
[4 5]
[1 9 3 4 5]
[1 9 3]
inclusive of first index, excl of second index:  [9 3]
```

### **STL - strings**

```go
	//? STL - strings

	greetings := "hello there friends!"
	fmt.Println(strings.Contains(greetings, "hello"))
	fmt.Println(strings.Contains(greetings, "apple"))

	fmt.Println(strings.ReplaceAll(greetings, "hello", "heloooo"))
	fmt.Println("original string:", greetings)
	fmt.Println("all upper letters: ", strings.ToUpper(greetings))

	fmt.Println(" ")

	fmt.Println("index of 'th' : ", strings.Index(greetings, "th"))

	fmt.Println("Split :", strings.Split(greetings, " "))

```
```
true
false
heloooo there friends!
original string: hello there friends!
all upper letters:  HELLO THERE FRIENDS!

index of 'th' :  6
Split : [hello there friends!]
```

### **STL - ints**

```go
	//? STL - ints?

	ages := []int{45, 56, 83, 26, 58, 38, 24, 23, 18}
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 26)
	fmt.Println(index)

	fmt.Println(" ")

	names := []string{"dog", "apple", "cat", "ball", "elephant"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "cat"))
```
```
[18 23 24 26 38 45 56 58 83]
3

[apple ball cat dog elephant]
2
```

### **LOOPS**
```go
	//? LOOPS
	x := 0

	for x < 5 {
		fmt.Println("value of x : ", x)
		x++
	}
```
```
value of x :  0
value of x :  1
value of x :  2
value of x :  3
value of x :  4
```

```go
	for i:= 0; i < 5; i++ {
		fmt.Println("value of i : ", i)
	}
```
```
value of i :  0
value of i :  1
value of i :  2
value of i :  3
value of i :  4
```
```go
	names := []string{"apple", "ball", "cat", "dog"}
	for i := 0; i < len(names); i++ {names[i])
	}
```

```
apple
ball
cat
dog
```

> better way

```go
	names := []string{"apple", "ball", "cat", "dog"}
	for index, value := range names {
		fmt.Println(index, ":", value)
	}
```
```
0 : apple
1 : ball
2 : cat
3 : dog
```

> if we dont wanna use index,we can replace it with _. same with the case of value

```go
	names := []string{"apple", "ball", "cat", "dog"}
	for _, value := range names {
		fmt.Println(value)
	}
```
```
apple
ball
cat
dog
```

> Value is just the local copy of the element in array  
```go
	names := []string{"apple", "ball", "cat", "dog"}
	for _, value := range names {
		fmt.Println(value)
		value = "new string"
	}
	fmt.Println(names)
```
```
apple
ball
cat
dog
[apple ball cat dog]
```


### **BOOLEANMS AND CONDITIONALS**
```go
	age := 25
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 55)
	fmt.Println(age != 55)
```
```
true
false
false
true
```

> conditionals
```go
	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is lesss than 40")
	} else {
		fmt.Println("age is greater than 40")
	}
```

```
age is less than 30
```


```go
	names := []string{"apple", "ball", "cat", "dog"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing on index 1 ")
			continue
		}

		fmt.Println(index, ".", value)

	}
```
```
0 . apple
continuing on index 1 
2 . cat
3 . dog
```

```go
	names := []string{"apple", "ball", "cat", "dog"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing on index 1 ")
			continue
		}
		if index > 2 {
			fmt.Println("breadking the loop", index)
			break
		}

		fmt.Println(index, ".", value)

	}
```
```
0 . apple
continuing on index 1 
2 . cat
breadking the loop 3
```

### **FUNCTIONS**

> func with return
```go
func sayGreetings (name string) string {
	// fmt.Println("good morning", name)
	return fmt.Sprintf("good morning %v", name)
}
```
>func without return
```go

func sayBye (name string){
	fmt.Println("good morning", name)
}
```

> SUPER FUNCS

> these take function as inputs, (function signatures)
```go
func cyclenames(names []string, f func(string) string) []string {
	var ans []string
	for _, val := range names {
		ans = append(ans, f(val))
	}
	return ans
}
```

```go
names := []string{"a", "b", "c", "d"}
fmt.Println(cyclenames(names, sayGreetings))
```


```go
func circleArea (r float64) float64 {
	return math.Pi * r * r
}

func main() {
	fmt.Println("area of circle of radius 3 : ", circleArea(3))
	fmt.Printf("area of circle rounded to 3 deciamls is %0.3f ", circleArea(3))
}
```
```
area of circle of radius 3 :  28.274333882308138
area of circle rounded to 3 deciamls is 28.274
```


### **MULTIPLE RETURN VALUES**

```go
package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	partsOfNames := strings.Split(s, " ")

	var initials []string
	for _, val := range partsOfNames {
		initials = append(initials, val[:1])
	}


	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	//? if there is no second name passed we return _ in place of initial
	return initials[0], "_" 

}

func main() {

	getInitials("harshith gandhe")
	fmt.Println(getInitials("harshith gandhe"))
	fmt.Println(getInitials("harshith"))
	fmt.Println(getInitials("harshith gandhe harry"))
	
	fn, sn := getInitials("gandhe harshith") 
	fmt.Println(fn, sn)
}
```

### **Package Scope**

> greetings.go
```go
package main

import "fmt"

func sayHello(n string) {
	fmt.Println("hello there, ",n)
}
```

> main.go
```go
package main

func main() {
	name := "harshith"

	sayHello(name)

}
```
```
hello there,  harshith
```

---
> main.go
```go
package main

var score1 = 10

func main() {

	var score2 = 20
	score2 = score2 + 0

	name := "harshith"

	sayHello(name)

	printScores()
	
}
```
> greetings.go
```go
package main

import "fmt"

func sayHello(n string) {
	fmt.Println("hello there, ",n)
}


func printScores() {
	fmt.Println(score1)
	fmt.Println(score2)
}
```

```
--> go run main.go greetings.go
# command-line-arguments
.\greetings.go:12:14: undefined: score2
```
> variables in main function cannot be seen by other packages, it has to be in the main package scope to be visible to the other packages

> Removing that score2 variable entirely from noth the packages and running the same would give us the same result



## **MAPS**

> All keys should have the same type, all values should also have the same type

```go
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println("the price of soup is ",menu["soup"])

	//? looping maps

	for k, v := range menu {
		fmt.Printf("item : %v -- price : %v \n", k, v)
	}

	fmt.Println(" ")

	phonebook := map[int]string{
		123456: "harshi",
		23456: "harshith",
		34567: "harshiii",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[123456])

	phonebook[123456] = "gandhe harshi"
	fmt.Println("phonebook[123456] : ",phonebook[123456])

	phonebook[23456] = "harsheyyyy"
	fmt.Println(phonebook)
```
```
map[pie:7.99 salad:6.99 soup:4.99 toffee pudding:3.55]
the price of soup is  4.99
item : pie -- price : 7.99
item : salad -- price : 6.99
item : toffee pudding -- price : 3.55
item : soup -- price : 4.99

map[23456:harshith 34567:harshiii 123456:harshi]
harshi
phonebook[123456] :  gandhe harshi
map[23456:harsheyyyy 34567:harshiii 123456:gandhe harshi]
```

> GO IS A PASS BY VALUE

> for group A types: strings, ints, bools, floats, arrays, structs

```go
package main

import "fmt"

func update(n string) {
	n = "b"
}

func main() {
	x := "a"
	update(x)
	fmt.Println(x)
}
```
```
a
```

To effect the variable in a function we have to return a value and set that variable to it

> To do that
```go
package main

import "fmt"

func update(n string) string {
	n = "b"
	return n
}

func main() {
	x := "a"
	x = update(x)
	fmt.Println(x)
}

```
```
b
```

> for group B types : slices, maps, functions


```go
func updateMenu (y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	menu := map[string]float64{
		"tea": 2.99,
		"milk": 1.99,
	}

	updateMenu(menu)

	fmt.Println(menu)
	fmt.Println("price of coffee:", menu["coffee"] )
}
```
```
map[coffee:2.99 milk:1.99 tea:2.99]
price of coffee: 2.99
```


## **POINTERS**


```go
package main

import "fmt"

func updateName(n string) {
	n = "updated"
}

func main() {

	temp := "old"

	updateName(temp)

	fmt.Println(temp)

	fmt.Println("address of the variable temp is", &temp)

	//? rteferencing the address of temp to m
	m := &temp

	//? dereferencing the address which should give us bavk the string
	temp2 := *m

	fmt.Println("temp2 is : ",temp2, "or", *m )

}
```
```
old
address of the variable temp is 0xc000052260
temp2 is :  old or old
```

> We can use pointers to pass (type A)variables between the functions

```go
package main

import "fmt"

func updateName(n string) {
	n = "updated"
}

func updateName2(n *string) {
	*n = "updated"
}

func main() {

	temp := "old"

	updateName(temp)

	fmt.Println(temp)

	fmt.Println("address of the variable temp is", &temp)

	//? rteferencing the address of temp to m
	m := &temp

	//? dereferencing the address which should give us bavk the string
	temp2 := *m

	fmt.Println("temp2 is : ",temp2, "or", *m )

	updateName2(m)

	fmt.Println("value of temp after updatename2 func is: ",temp)

}
```
```
old
address of the variable temp is 0xc000052260
temp2 is :  old or old
value of temp after updatename2 func is:  updated
```


## **STRUCTS**


> bill.go
```go
package main

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
```
>main.go
```go
package main

import "fmt"


func main() {
	mybill := newBill("harshi's bill")
	fmt.Printf("mybill: %v\n", mybill)

}
```
```
go run .\main.go .\bill.go
mybill: {harshi's bill map[] 0}
```

### **Reciver functions**
> these functions can only be called on the objects that are associated with the func
>in below code block, b bill represents the reciver object and b is that object that will be used in the function block

```go
func (b bill) format () {
	
}
```

> FULL EXAMPLE	

> main.go
```go
package main

import "fmt"


func main() {
	mybill := newBill("harshi's bill")
	fmt.Printf("mybill: %v\n", mybill)

	fmt.Print("formatted Bill is: \n", mybill.format())
}
```
> bill.go
```go
package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// ? create a new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}
	return b
}

// ? format the bill
func (b bill) format() string {
	fs := "bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v...$%v \n", k+":", v)
		total += v
	}

	//? total
	fs += fmt.Sprintf("%-25v...$%v", "total: ", total) //? that -25 is for spacing and stuff, change it to 25 and see diff
	return fs
}
```
```
mybill: {harshi's bill map[cake:3.99 pie:5.99] 0}
formatted Bill is:
bill breakdown:
cake:                    ...$3.99
pie:                     ...$5.99
total:                   ...$9.98
```


> FINAL CODE

> bill.go
```go
package main

import "fmt"

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
```
> main.go
```go
package main

import "fmt"

func main() {
	mybill := newBill("harshi's bill")
	// fmt.Printf("mybill: %v\n", mybill)

	mybill.addItem("onion soup", 4.99)
	mybill.addItem("maggi", 3.99)
	mybill.addItem("wine", 6.99)

	mybill.updateTip(10)
	
	fmt.Print(mybill.format())
}
```
> output
```
bill breakdown: 
wine:                    ...$6.99
onion soup:              ...$4.99
maggi:                   ...$3.99
tip:                     ...$10
total:                   ...$25.97
```

## **USER INPUT**

> main.go
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createBill () bill {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Create a new bill name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	b := newBill(name)
	fmt.Println("created bill - ", b.name)

	return b
}


func main() {
	mybill := createBill()
	fmt.Println(mybill)
}

```
> and bill.go
```go
package main

import "fmt"

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
```
> will create a bill by taking name from the user input

> we cannot really create new reader and thae corresponding code everytime, we can just create a helper function

> This is how it should look like

>main.go

```go
package main

import (
	"bufio"
	"fmt"
	"os"
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


func main() {
	mybill := createBill()
	fmt.Println(mybill)
}

```

> Adding saving func and some refactoring the code

> main.go
```go
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
```
> bill.go
```go
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
```

## **INTERFACES**

> Group toagther the types(user defined) which have same functions (like circle and square are data types and area and perimeter are similar functions) we can group them togather and can create a func which takes anything from the group as input and does the functions (which have same signature btw) in the function block

> main.go
```go
package main

import (
	"fmt"
	"math"
)

// shape interface
type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}
type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
}

func main() {
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
}
```