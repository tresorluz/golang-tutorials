package main

import "fmt"

func main() {
	// Variable declaration
	age := 33
	ageTwo := "34"
	name := "Tresor"

	// This Print statement doesn't add a newline
	fmt.Print("Hello, ")
	fmt.Print("World! \n")

	// This Print statement add a newline automatically
	fmt.Println("Hello world")
	// Inject variables in print statement
	fmt.Println("My name is", name, "and I'm", age, "Years old")

	// Formatted strings %_ format specifier
	//%v for variable %q to get quotes around string %T for data type
	fmt.Printf("My name is %v and I'm %v \n", name, age)
	fmt.Printf("My name is %q and I'm %v \n", name, age)
	fmt.Printf("My name is %T and I'm %T \n", name, age)
	fmt.Printf("Your score is %f points \n", 255.55)
	fmt.Printf("Your score is %0.1f \n", 255.55)

	//sprintf save formatted strings
	var str = fmt.Sprintf("My name is %v and I'm %v \n", name, ageTwo)
	fmt.Println("The saved string is: ",str)
}