package main

import "fmt"

func main() {
	//normal construct with all the 3 expression
	for i, sum := 0, 0; i < 10; i++ {
		sum = sum + i
		fmt.Println(sum)
	}
	println("###")
	//construct with some of the missing expressions
	for i := 0; ; i++ {
		println(i)
		if i > 5 {
			break
		}
	}
	println("###")
	//while like construct
	isGreaterthan10 := false
	num := 0
	for isGreaterthan10 == false {
		num = num + 1
		println(num)
		if num > 10 {
			fmt.Println("counter past 10")
			isGreaterthan10 = true
		}
	}
	println("###")
	//infinite loop construct
	sum := 0
	for {
		sum = sum + 1
		if sum > 10 {
			fmt.Println("sum past 10")
			break
		}
	}

}
