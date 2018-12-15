package main

import "fmt"
import "time"
import "math/rand"

func main() {
	rand.Seed(time.Now().UnixNano())
	//go's for loop doesn't allow
	//var keyword so shorthand of variable declaration has to be used
	for i := 0; i < 10; i++ {
		var countryNumber int = rand.Intn(3)
		//switch without an condition and one single initializer
		switch name := "Sujal"; {
		case countryNumber == 0:
			fmt.Println("Namaste", name)
		case countryNumber == 1:
			fmt.Println("Ohio", name)
		case countryNumber == 2:
			fmt.Println("Hello", name)
		case countryNumber == 3:
			fmt.Println("Merhaba", name)
		}
	}
}
