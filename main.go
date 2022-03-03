package main

import (
	"fmt"
	"math"
)


func sayGreeting(name string){
	fmt.Printf("Good morning %v \n", name)
}
func sayBye(name string){
	fmt.Printf("Goodbye %v \n", name)
}
func cycleNames(n []string, f func(string)){
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64{
	return math.Pi * r * r
}
func main() {
	names := []string{"Tresor", "Victoria", "Elyon", "Luzingamu"}

	sayGreeting("Tresor")
	sayGreeting("Victoria")
	sayBye("Elyon")
	cycleNames(names, sayGreeting)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
}