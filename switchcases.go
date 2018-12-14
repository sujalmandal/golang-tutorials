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
	switch x {
	case 1:
		println("rand number is ", x)
	case 2:
		println("rand number is ", x)
	case 3:
		println("rand number is ", x)
	}
}
