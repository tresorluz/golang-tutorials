package main

import "fmt"

func main() {
	// declare a string
	var nameOne string = "Elyon"
	var nameTwo = "Victoria"
	var namethree string

	fmt.Println(nameOne, nameTwo)
	nameOne = "Ely"
	namethree = "Tresor"

	fmt.Println(nameOne, nameTwo, namethree)
	// shorthand string declaration can't be used outside a function
	nameFour := "Luzingamu"
	fmt.Println(nameFour)

	// declare an int
	var ageOne int = 4
	var ageTwo = 27
	ageThree := 33

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits(8,16,32,64) and memory
	// int8 bit size is 8
	var numOne int8 = 25
	// unsingned int cannot be a negative number
	var numTwo uint = 3

	// declare a float needs to declare a bit size
	var scoreOne float32 = 25.98
	// can be a very large number
	var scoreTwo float64 = 78654321890765.3452

	fmt.Println(numOne, numTwo, scoreOne, scoreTwo)

}