package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func main() {
	//go's if loop & switch can have initializer expression
	if os := runtime.GOOS; os != "windows" {
		fmt.Println(os)
	} else { // else cannot begin in the next line
		fmt.Println(os)
	}
	x := rand.Intn(3)
	//go's switch only executes one case & doesn't need a break statement
	//this is an example of a switch without default
	switch x {
	case 1:
		fmt.Println("rand number is ", x)
	case 2:
		fmt.Println("rand number is ", x)
	case 3:
		fmt.Println("rand number is ", x)
	}
	//example of switch with default
	x = 100
	switch x {
	case 1:
		fmt.Println("rand number is ", x)
	case 2:
		fmt.Println("rand number is ", x)
	case 3:
		fmt.Println("rand number is ", x)
	default:
		fmt.Println("default case, none of the cases matched")
	}
}
