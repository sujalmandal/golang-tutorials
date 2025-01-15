package main

import "fmt"

func main() {
	var myMap map[string]string = map[string]string{
		"Name": "John",
		"Age":  "25",
	}

	for k, v := range myMap {
		fmt.Println(k, " : ", v)
	}

	//updating a map
	myMap["Age"] = "32"
	fmt.Println("Updated Map : ", myMap)

	delete(myMap, "Age")

	fmt.Println("Map after key deleted : ", myMap)

	var age string = myMap["Age"]

	fmt.Printf("Age : [%v]", age)
}
