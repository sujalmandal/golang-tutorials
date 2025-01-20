package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Enter something to save: ")
	data, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	data = strings.TrimSpace(data)
	dataBytes := []byte(data)
	err := os.WriteFile("../temp/files-example.txt", dataBytes, 0777)

	if err == nil {
		fmt.Println("File written")
	} else {
		errMsg := err.Error()
		fmt.Printf("failed to write file! %v", errMsg)
	}
}
