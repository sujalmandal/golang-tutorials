package main

func main() {
	var p Person = newPerson("Sujal", "Mandal", 32)
	var addr Address = newAddress("Dubai", "United Arab Emirates", "Dubai")
	//original person object
	p.print("person instance")
	//update using func that sends a copy
	copyOfPerson := p.getUpdatedPersonWithAddress(addr)
	p.print("person instance after calling getUpdatedPersonWithAddress()")
	copyOfPerson.print("copy of person instance returned by getUpdatedPersonWithAddress()")
	//mutate using func that takes pointer of person
	p.updatePersonAddress(addr)
	p.print("mutated person instance after calling updatePersonAddress()")
}

// receiver function on Person type - will not mutate the person instance
func (p Person) getUpdatedPersonWithAddress(address Address) Person {
	//receives a person instance copy to work with
	p.address = address
	return p
}

// receiver function on *Person (Pointer) type - will mutate the person instance
func (p *Person) updatePersonAddress(address Address) {
	//receives a person pointer instance copy - that points to the original instance
	//thus, any operation done on this pointer will end up mutating the original person instance
	p.address = address
}
