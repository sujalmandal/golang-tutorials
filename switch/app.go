package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string) string {
	fmt.Print(prompt, ": ")
	bufReader := bufio.NewReader(os.Stdin)
	input, _ := bufReader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	i := getInput("Enter a number")
	num, err := strconv.ParseInt(i, 10, 32)
	if err != nil {
		fmt.Printf("You did not enter a valid value! [%v]\n", i)
	}

	switch num {
	case 10:
		fmt.Println("You entered 10")
	case 100:
		fmt.Println("You entered 100")
	case 1000:
		fmt.Println("You entered 1000")
	case 10_000:
		fmt.Println("You entered 10,000")
	case 100_000:
		fmt.Println("You entered 100,000")
	default:
		fmt.Println("You entered some other number!")
	}
}
