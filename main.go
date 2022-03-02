package main

import "fmt"

func main() {
	//Arrays declaration
	var agesOne[3]int = [3]int{20,25,30}
	var agesTwo = [3]int{20,25,30}
	fmt.Println(agesOne, agesTwo, len(agesOne))

	names := [4]string{"Tresor", "Victoria", "Elyon", "Luzingamu"} 
	names[2] = "Ely"
	fmt.Println(names, len(names))

	// slice uses arrays under the hood
	var scores = []int{100, 50, 60}
	scores[2] = 25

	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[1:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

}