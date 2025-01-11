package main

import "fmt"

func main() {
	// Declaring an array of strings - 3 elements, initialized with default values
	var names [3]string
	fmt.Println(names)
	// Declare & initialize an array of integers
	var ages [3]int = [3]int{28, 31, 35}
	fmt.Println(ages)
	//Short hand for declaring & initializing an array
	var designations = [3]string{"Junior Developer", "Senior Developer", "System Developer"}
	fmt.Println("All positions held are -> ", designations)
	// for length of the array, use the inbuilt function len()
	fmt.Println("Total positions held ->", len(designations))

	// Accessing elements of an array
	fmt.Println("First position held ->", designations[0])
	//update element of an existing array
	designations[0] = "Java Developer"
	fmt.Println("First position held ->", designations[0])

	// this will throw an error as the array is of length 3 - will throw out of bounds error
	//designations[4] = "Python Developer"

	// slices -> a slice is a reference to a contiguous segment of an array
	// slices can be dynamically resized

	var companiesWorkedAt = []string{"Mindtree", "TimeInc", "EMC2", "OpenText"}
	fmt.Println("Companies worked at ->", companiesWorkedAt)

	// append to a slice - this will create a new slice
	var updatedCompanies = append(companiesWorkedAt, "Operative")
	fmt.Println("Updated companies worked at ->", updatedCompanies)
	//append operation does not modify the original slice
	fmt.Println("Companies worked at previously ->", companiesWorkedAt)

	//slice ranges
	// [start: end] - start is inclusive, end is exclusive
	var first3Companies = updatedCompanies[:3]
	//slice has total 5 items
	// "Mindtree", "TimeInc", "EMC2", "OpenText", "Operative"
	// last 2 companies
	lastTwoCompanies := updatedCompanies[2:]
	// all companies except the first and last
	allCompaniesExceptFirstAndLast := updatedCompanies[1:4]

	fmt.Println("First 3 companies ->", first3Companies)
	fmt.Println("Last 2 companies ->", lastTwoCompanies)
	fmt.Println("All companies except first and last ->", allCompaniesExceptFirstAndLast)
}
