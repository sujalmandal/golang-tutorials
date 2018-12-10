package main

import "fmt"

func main() {
	fmt.Println(nextPairOfFibonacci(3, 5))
}

func nextPairOfFibonacci(num1, num2 int) (next1, next2 int) {
	return num1 + num2, num1 + num2 + num2
}
