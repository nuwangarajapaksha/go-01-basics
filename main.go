package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, Nuwa")

	//strings
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string
	nameFour := "name4"
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	nameOne = "n1"
	nameTwo = "n2"
	nameThree = "n3"
	nameFour = "n4"
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	//ints
	var ageOne int = 1
	var ageTwo = 2
	var ageThree int
	ageFour := 4
	fmt.Println(ageOne, ageTwo, ageThree, ageFour)

	//bits & memory
	//var numOne int8 = 25
	//var numTwo int8 = -129
	//var numThree uint8 = 25

	//floats
	var scoreOne float32 = 25.98
	var scoroTwo float64 = 3094234374298374.3
	scoroThree := 1.5
	fmt.Println(scoreOne, scoroTwo, scoroThree)

	//Print
	fmt.Print("Hello, ")
	fmt.Print("World! \n")
	fmt.Print("new line \n")

	//Println
	age := 35
	name := "shaun"
	fmt.Println("hellow ninjas!")
	fmt.Println("goodbye ninjas!")
	fmt.Println("my age is", age, "and my name is", name)

	//Printf (formatted string ) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age if of type %T \n", age)
	fmt.Printf("you scored %f points! \n", 225.55)
	fmt.Printf("you scored %0.1f points! \n", 225.55)

	//Sprintf (save formatted string)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)

	//arrys
	var ages [3]int = [3]int{20, 25, 30}
	var ages2 = [3]int{20, 25, 30}
	ages3 := [3]int{20, 25, 30}
	names := [5]string{"yoshi", "mario", "peach", "bower", "lugi"}
	fmt.Println(ages, ages2, ages3, names)

	//slices (use arrys under the hood)
	var scores = []int{100, 200, 300}
	scores[2] = 25
	scores = append(scores, 50)
	fmt.Println(scores)

	//slice range
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)
	rangeOne = append(rangeOne, "kopa")
	fmt.Println(rangeOne)

	//The Standard Library
	//strings package
	greeting := "Hello! there friends"
	fmt.Println(strings.Contains(greeting, "Hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello!", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "th"))
	fmt.Println(strings.Split(greeting, " "))

	//sort packege
	ages5 := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages5)
	fmt.Println(ages)

	index := sort.SearchInts(ages5, 30)
	fmt.Println(index)

	names2 := []string{"yoshi", "mario", "peach", "bower", "lugi"}
	sort.Strings(names2)
	fmt.Println(names2)
	index2 := sort.SearchStrings(names2, "bower")
	fmt.Println(index2)

	//for loops

	//normal for loop
	x := 0
	for x < 5 {
		fmt.Println("value of x: ", x)
		x++
	}

	//indexed for loop
	for i := 0; i < 5; i++ {
		fmt.Println("value of i: ", i)
	}

	names3 := []string{"yoshi", "mario", "peach", "bower", "lugi"}
	for i := 0; i < len(names3); i++ {
		fmt.Println(names3[i])
	}

	//range for loop
	for index, value := range names3 {
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

	for _, value := range names3 {
		fmt.Printf("the value is %v \n", value)
	}

	//booleans and conditionals
	age2 := 45

	fmt.Println(age2 <= 50)
	fmt.Println(age2 >= 50)
	fmt.Println(age2 == 50)
	fmt.Println(age2 != 50)

	if age2 < 30 {
		fmt.Println("age is less than 30")
	} else if age2 < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}

	names4 := []string{"yoshi", "mario", "peach", "bower", "lugi"}
	for i, i2 := range names4 {
		if i == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if i > 2 {
			fmt.Println("braking at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v \n", i, i2)
	}

	//functions
	sayGreeting("mario")
	sayGreeting("lugi")
	sayBay("mario")
	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	cycleNames([]string{"cloud", "tifa", "barret"}, sayBay)

	a1 := cycleArea(1)
	a2 := cycleArea(5)
	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f \n", a1, a2)

	//Multiple Return Values
	fn1, sn1 := getInitials("tifa lockhart")
	fmt.Println(fn1, sn1)
	fn2, sn2 := getInitials("tifa")
	fmt.Println(fn2, sn2)

	//package scope
	sayHello("mario")
	for _, i2 := range points {
		fmt.Println(i2)
	}
	showScore()

	//maps
	//strings as key type
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//ints as key type
	phonebook := map[int]string{
		267584967: "mario",
		984759373: "luigi",
		845775485: "peach",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[267584967])
	phonebook[984759373] = "bowser"
	fmt.Println(phonebook)

	//pass by value
	// group A types -> strings,ints,boole,flots,arrays,structs
	name1 := "tifa"
	name1 = updateName(name1)
	fmt.Println(name1)
	//group A types -> slices,maps, functions
	menu1 := map[string]float64{
		"pie":       5.95,
		"ice creme": 3.99,
	}
	updateMenu(menu1)
	fmt.Println(menu1)

	//pointers
	name2 := "tifa"
	fmt.Println("memory address of name is:", &name2)
	m := &name2
	fmt.Println("memory address:", m)
	fmt.Println("value at memory address", *m)
	updateName2(m)
	fmt.Println(name2)

	//struct & custom types
	mybill := newBill("marios bill")
	fmt.Println(mybill)

	//receiver functions
	fmt.Println(mybill.format())

	//Reciver functions with pointers
	mybill.addItem("onion soup", 4.50)
	mybill.addItem("veg pie", 8.95)
	mybill.addItem("toffee pudding", 4.95)
	mybill.addItem("coffee", 3.25)
	mybill.updateTip(10)
	fmt.Println(mybill.format())

	//user input
	myBill2 := createBill()
	fmt.Println(myBill2)
	promptOptions(myBill2)

}

func sayGreeting(n string) {
	fmt.Printf("Good Morning %v \n", n)
}

func sayBay(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, i2 := range n {
		f(i2)
	}
}

func cycleArea(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	fmt.Println(names)

	var initials []string
	for _, i2 := range names {
		fmt.Println(i2[:1])
		initials = append(initials, i2[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"

}

// package scope
var score = 99.5

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
	y["pie"] = 2.99
}

func updateName2(x *string) {
	*x = "wedge"
}

// user input
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("choose option (a - add item, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)

	//switch statement
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item Price", reader)
		fmt.Println("you chose a")

		//parsing floats
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("item added: ", name, ":", price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($) ", reader)
		fmt.Println("you chose t")
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip added: ", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you chose to save the bill", b)
		fmt.Println("you saved the file - ", b.name)
	default:
		fmt.Println("that was not a valid option")
		promptOptions(b)
	}
}
