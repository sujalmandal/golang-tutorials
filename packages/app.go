package main

// package level declaration
var logLevel = 0

// to run this example use the folowing command:
// go run app.go logger.go
func main() {
	// uses a function defined in the logger.go package which is also defined
	// in the main package - same package as this file, granting this file access
	// to all functions and variables defined in the logger.go file

	logInfo("Hello, World!")
}
