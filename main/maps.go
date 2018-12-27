package main

import (
	"fmt"
)

type Movie struct {
	releaseYear  string
	directorName string
	cast         []string
}

var movieInfo map[string]Movie

func main() {
	//maps in go start with a special keyword "map" followed by the type of index inside square braces
	//and the type of key e.g. myMap:=map[string]string
	var empContactDb map[string]string
	empContactDb = make(map[string]string)
	//put values inside the map
	empContactDb["sujal"] = "8904711202"
	empContactDb["david"] = "9943103892"
	empContactDb["david"] = "8904711202"
	//access value inside the map by key
	fmt.Println(empContactDb["sujal"])
	//map definitions with mao literals
	empAddresses := map[string]string{"Sujal": "12th street, uguabuga colony, ouas residency 2nd Floor, Rno : 089"}
	fmt.Println(empAddresses)
	//map with struct as key
	movieInfo = map[string]Movie{"aquaman": Movie{releaseYear: "2018", directorName: "James Wan", cast: []string{"Jason Momoa", "Amber Heard"}}}
	fmt.Println(movieInfo)
}
