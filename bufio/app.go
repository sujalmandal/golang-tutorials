package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter a text -> ")
	//represents the console for inputs
	in := bufio.NewReader(os.Stdin)
	input, _ := in.ReadString('\n')

	out := bufio.NewWriter(os.Stdout)

	totalBytesWritten, _ := out.WriteString(fmt.Sprintf("You entered %v", input))
	//bufio won't flush data to the underlying file (Stdout) immediately, use flush() to force it!
	out.Flush()

	fmt.Println("total writes written :", totalBytesWritten)
}
