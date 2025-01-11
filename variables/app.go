// package must be main for the main function to be the entry point of the program
package main

//system package that is used to print to the console
import "fmt"

// this is a main function - this is the entry point of the program
func main() {
	//full variable declaration (initialized with default value for the variable type)
	var name string
	//variable declaration with initial value
	var age int = 25

	fmt.Println("Name is: ", name)
	fmt.Println("Age is: ", age)

	//short variable declaration
	middleName := "Kumar"
	fmt.Println("Middle Name is: ", middleName)

	// many different kinds of variable types
	// int, float64, string, bool, byte

	var intExample int = 1
	var floatExample float64 = 1.1
	var boolExample bool = true
	// single quote is used for byte type and double quote is used for string types
	var stringExample string = "Hello"
	var byteExample byte = 'a'

	fmt.Println("intExample: ", intExample)
	fmt.Println("floatExample: ", floatExample)
	fmt.Println("boolExample: ", boolExample)
	fmt.Println("stringExample: ", stringExample)
	fmt.Println("byteExample: ", byteExample)

	// variables can be declared with signed or unsigned types & we can
	// specify the size of the variable type

	// signed int types (int8, int16, int32, int64)
	var int8Example int8 = -128 //total values -> 2^8 = 256, [-128 to 127]
	fmt.Println("int8Example: ", int8Example)
	int8Example = 127
	fmt.Println("int8Example: ", int8Example)

}
