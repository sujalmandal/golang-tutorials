package main

import "fmt"

func main() {
	// a pointer of string type
	var namePtr *string
	var name string = "Sujal Mandal"
	namePtr = &name

	fmt.Printf("memory address of '%v' is %v\n", name, namePtr)

	nameUpdate(name)
	fmt.Printf("name -> [%v] after nameUpdate()\n", name)

	namePtrUpdate(namePtr)
	fmt.Printf("name -> [%v] after namePtrUpdate()\n", name)
}

func nameUpdate(name string) {
	name = name + " - updated"
}

func namePtrUpdate(namePtr *string) {
	*namePtr = *namePtr + " - updated"
}
