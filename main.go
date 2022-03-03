package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//looping through a map
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		267584967: "Tresor",
		984759373: "Victoria",
		845775485: "Elyon",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[267584967])
}