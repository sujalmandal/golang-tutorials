package main

import "fmt"

//ENGLISH skills
const ENGLISH = "ENGLISH"

//FRENCH skills
const FRENCH = "FRENCH"

//TURKISH skills
const TURKISH = "TURKISH"

//GERMAN skills
const GERMAN = "GERMAN"

//HINDI skills
const HINDI = "HINDI"

//Employee struct defined outside main
type Employee struct {
	name              string
	age               int
	citizenShipStatus bool
	//example of nested struct
	englishSkill LangSkill
	//example of nested array(size 2) of struct
	otherLangSkill [2]LangSkill
}

//LangSkill another struct
type LangSkill struct {
	skill       string
	proficiency int
}

func main() {
	//struct definition by ommiting some fields
	emp1 := Employee{name: "Sujal", age: 25, citizenShipStatus: false}
	fmt.Println(emp1)
	//full definition with anynomous values
	emp2 := Employee{"Busra", 24, false, LangSkill{ENGLISH, 7}, [2]LangSkill{{TURKISH, 10}, {GERMAN, 5}}}
	fmt.Println(emp2)
	//standalone LangSKill struct declaration
	engSkillsEmp1 := LangSkill{ENGLISH, 9}
	otherSkillsEmp1 := [2]LangSkill{{HINDI, 9}, {TURKISH, 3}}
	//setting using dot operator
	emp1.englishSkill = engSkillsEmp1
	emp1.otherLangSkill = otherSkillsEmp1
	fmt.Println(emp1)
	//storing a slice of structs
	employees := []Employee{emp1, emp2}
	fmt.Println("All employee info: ", employees)
}
