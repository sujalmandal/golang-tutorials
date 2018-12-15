package main

import "fmt"

func main() {
	//defer functions don't run immediately instead they are
	//pused to the function stack & start executing after the function
	//who spawned the defered functions has finished execution
	defer fmt.Println("\ngo!!")
	for i := 1; i < 6; i++ {
		defer fmt.Print(i, "..")
	}
	defer fmt.Println("we shall start in ")
	defer fmt.Println("Sujal")
	fmt.Print("Hello, call me ")
	//deferred functions are executed last in first out manner
	//therefore they can be used for reversing things
}
