package main

import "fmt"

type Person struct {
	firstName  string
	middleName string
	lastName   string
	age        int
	address    Address
}

func newPerson(firstName string, lastName string, age int) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

func (p *Person) print(msg string) {
	fmt.Printf("------%v------", msg)
	fmt.Printf(
		"\nName -> %v\nAge -> %v\nAddress ->\n%v\n",
		(p.firstName + " " + p.middleName + " " + p.lastName),
		p.age, p.address.toString(),
	)
}

type Address struct {
	addressLine1 string
	addressLine2 string
	addressLine3 string
	country      string
	state        string
	postalCode   int
}

func newAddress(line1 string, country string, state string) Address {
	return Address{
		addressLine1: line1,
		country:      country,
		state:        state,
	}
}

func (a Address) toString() string {
	return fmt.Sprintf("1.%-40v\n2.%-40v\n3.%-40v\n4.%-40v\n5.%-40v\n6.%-40v",
		fmt.Sprintf("Line 1   		: %v", a.addressLine1),
		fmt.Sprintf("Line 2   		: %v", a.addressLine2),
		fmt.Sprintf("Line 3   		: %v", a.addressLine3),
		fmt.Sprintf("Country  		: %v", a.country),
		fmt.Sprintf("State    		: %v", a.state),
		fmt.Sprintf("PostalCode    	: %v", a.postalCode),
	)
}
