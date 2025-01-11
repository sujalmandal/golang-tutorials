package main

import "fmt"

func main() {
	count := 0
	// go only has 'for' keyword for loops
	fmt.Println("*******************************************")
	//variation 1
	// for loop that looks like a while loop in other languages
	for count < 10 {
		fmt.Printf("[variation 1] - Counted until %v\n", count)
		count++
	}

	//variation 2
	// for loop that looks like a for loop in other languages
	fmt.Println("*******************************************")
	// same line can have multiple statements separated by ;
	for count = 0; count < 10; count++ {
		fmt.Printf("[variation 2] - Counted until %v\n", count)
	}
	fmt.Println("*******************************************")

	//variation 3 - applicable on arrays, slices, etc

	var names = [3]string{"sujal", "busra"}
	for index, item := range names {
		fmt.Printf("Item %q is present at the index %v\n", item, index)
	}

	// go through all the values without using the index

	for _, item := range names {
		fmt.Printf("Item %q is present\n", item)
	}
}
