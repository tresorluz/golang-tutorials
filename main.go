package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greetings := "Hello there friends!"
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	fmt.Println(strings.Contains(greetings, "Hello"))
	fmt.Println((strings.ReplaceAll(greetings, "Hello", "Hi")))
	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Index(greetings, "ll"))
	fmt.Println(strings.Split(greetings, " "))

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"Tresor", "Victoria", "Elyon", "Luzingamu"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Victoria"))
}