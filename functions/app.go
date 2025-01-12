package main

import "fmt"

func main() {
	//invoke a function
	greet("Hello Sujal!")
	greetAll([]string{"Sujal", "Busra", "Leila"}, greet)
	fmt.Println("Sum of 3 and 2 is: ", add(3, 2))
	quotient, remainder := divide(3, 2)
	fmt.Printf("Dividing 3 by 2 gives quotient: %v & remainder %v", quotient, remainder)
}

// functions can take other functions as params
func greetAll(names []string, greeter func(string)) {
	for i := 0; i < len(names); i++ {
		greeter(names[i])
	}
}

// define a function with parameters
func greet(name string) {
	fmt.Println(name)
}

// functions returning a value
func add(num1 int, num2 int) int {
	return num1 + num2
}

// functions can return multiple values
func divide(num1 int, num2 int) (int, int) {
	var quotient = num1 / num2
	var remainder = num1 % num2
	return quotient, remainder
}
