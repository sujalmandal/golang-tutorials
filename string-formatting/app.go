package main

import "fmt"

func main() {
	// prints "Hello, World!" and a newline
	fmt.Println("Hello, World!")
	//only prints "Hello, World!"
	fmt.Print("Hello, World!")
	fmt.Print("How is it going?")

	//variables can be used in Print functions
	var name string = "Sujal"
	lName := "Mandal"
	//variables and strings can be interpolated
	fmt.Println("Hello, I am", name, ", and my last name is", lName)

	//formatting strings
	fmt.Printf("Hello, I am %s, and my last name is %s\n", name, lName)

	//we used format specifiers
	// %v -> for default
	// %s -> for strings
	// %d -> for integers

	age := 31
	fmt.Printf("%v's age is %v\n", name, age)

	//special format specifiers
	// %q -> for double-quoted strings
	// %T -> for type of the value
	// %t -> for boolean values
	// %f -> for floating-point numbers, for controlling precision use %0.xf where x is the number of decimal places

	fmt.Printf("Type of name is %T, age is of type %T\n", name, age)
	fmt.Printf("%v's family name is %q\n", name, lName)
	fmt.Printf("Is %v an adult? -> %t\n", name, age >= 18)

	var monthlyExpenses float64 = 8029.121821
	fmt.Printf("%v's monthly expenses are %f, but we can round it off to %0.2f\n", name, monthlyExpenses, monthlyExpenses)

	//instead of using fmt.Printf, we can use Sprintf to format strings and store them in a variable

	var formattedString = fmt.Sprintf("SELECT * from users where name = '%s'", name)
	fmt.Println("Generated SQL is -> [", formattedString, "]")

	// other format specifiers can include the below
	// %x -> for hexadecimal strings
	// %b -> for binary strings
	// %o -> for octal strings

	var decimalNumber int = 1024
	var numberInHex = fmt.Sprintf("%x", decimalNumber)
	fmt.Println("Decimal number", decimalNumber, "in hexadecimal is", numberInHex)

	var numberInBinary = fmt.Sprintf("%b", decimalNumber)
	fmt.Println("Decimal number", decimalNumber, "in binary is", numberInBinary)

	var numberInOctal = fmt.Sprintf("%o", decimalNumber)
	fmt.Println("Decimal number", decimalNumber, "in octal is", numberInOctal)
}
