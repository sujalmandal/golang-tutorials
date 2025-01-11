package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string = "Sujal Mandalarian"
	//check if a substring is contained in a string
	fmt.Printf("Is %q a Mandal? : %t\n", name, strings.Contains(name, "Mandal"))
	//replace a substring with another substring
	updatedName := strings.ReplaceAll(name, "Mandalarian", "Mandal")
	fmt.Printf("Updated Name: %q\n", updatedName)

	fmt.Printf("%q to uppercase ->%q\n", updatedName, strings.ToUpper(updatedName))

	namesAsCsv := "Sujal Mandal, Busra Ercelik, Shankha Jana, Jhon Doe"

	var names = strings.Split(namesAsCsv, ",")
	fmt.Printf("Names: %v\n", names)

	var positionOfSujal = strings.Index(name, "uj")
	fmt.Println("Position of Sujal in name: ", positionOfSujal)
}
