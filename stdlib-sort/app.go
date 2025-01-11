package main

import (
	"fmt"
	"sort"
)

func main() {
	//sorting ints
	var numbers = []int{2, 3, 1, 4, 5}
	fmt.Println("Before sorting", numbers)
	sort.Ints(numbers)
	fmt.Println("After sorting", numbers)

	//searching for a number
	var index = sort.SearchInts(numbers, 4)
	fmt.Println("Index of 4 is", index)
	// if the item was not found, the len of the slice is returned instead
	fmt.Println("Index of 100 is", sort.SearchInts(numbers, 100))

	var names = []string{"john", "wick", "ethan", "hunt"}
	fmt.Println("Before sorting", names)
	sort.Strings(names)
	fmt.Println("After sorting", names)

	fmt.Println("Index of john is", sort.SearchStrings(names, "john"))
}
