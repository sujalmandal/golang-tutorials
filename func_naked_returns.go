package main

import "fmt"

func main() {
	var result = add(10, 100)
	fmt.Printf("Sum is : %v", result)
}

//example of naked return, use only in short methods which are too obvious
func add(num1 int, num2 int) (sum int) {
	sum = num1 + num2
	return
}
