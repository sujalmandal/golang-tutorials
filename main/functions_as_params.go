package main

import "fmt"

//an function that takes another function as an argument
func compute(operation func(int, int) int, num1 int, num2 int) int {
	return operation(num1, num2)
}

func main() {
	//define the behavior that calculates the area given the length & breadth
	area := func(length int, breadth int) int {
		return length * breadth
	}
	//call the method compute while passing a behavior & the values it has to operate on
	fmt.Println(compute(area, 10, 20))
}
