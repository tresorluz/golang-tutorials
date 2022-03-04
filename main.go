package main

import "fmt"

func main() {
	mybill := newBill("Tresor's bill")
	
	fmt.Println(mybill.format())
}