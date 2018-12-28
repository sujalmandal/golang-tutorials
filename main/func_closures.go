package main

import "fmt"

func main() {
	incrementer := func(incrementBy int) func() int {
		//default value is 0
		initialNumber := 0
		return func() int {
			initialNumber += incrementBy
			return initialNumber
		}
	}
	//call the main closure function to get function instances which are bound to "initialNumbers"
	incrementByOne := incrementer(1)
	incrementByTwo := incrementer(2)
	for index := 0; index < 10; index++ {
		//each call remembers the initial state
		fmt.Println(incrementByOne(), incrementByOne())
		fmt.Println(incrementByTwo(), incrementByTwo())
	}
}
