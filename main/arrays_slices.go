package main

import "fmt"

func main() {
	//empty array declaration
	var movies [3]string
	fmt.Println(movies)
	//array declared with size defined
	var animes [3]string
	fmt.Println(animes)
	//putting values inside the arrays defined above
	//arrays have fixed size & cannot be resized, the exact size has to be supplied
	movies = [3]string{"Trancendence", "Prometheus", "Europa Report"}
	fmt.Println(movies)
	//slices are references to arrays & can be dynamically resized
	//an slice is made from an array by specifying the start & end window within []
	mySlice := movies[0:3]
	fmt.Println(mySlice)
	//slice declaration happens using empty braces e.g. []
	//no lenght is specified
	var myEmptySlice []int
	fmt.Println(myEmptySlice)
	//slices can modify the original array they were made from
	mySlice[0] = mySlice[0] + ":2"
	fmt.Println(movies)
	//slices can be fully used to operate anything that can be operated by arrays
	fibonacciNumbers := []int{0, 1, 0, 0, 0, 0, 0}
	for index := 0; ; index++ {
		index1 := index
		index2 := index + 1
		if index2 > len(fibonacciNumbers)-2 {
			break
		}
		num1 := fibonacciNumbers[index1]
		num2 := fibonacciNumbers[index2]
		index3 := index2 + 1
		num3 := num1 + num2
		fibonacciNumbers[index3] = num3
	}
	fmt.Println(fibonacciNumbers)
	//double dimension slices
	var matrix [][]int = [][]int{[]int{1, 2, 3},
		[]int{4, 5, 6}}
	fmt.Println(matrix)
	//cap() defines the total size of the underlying array & len() defines the total size of the slice window
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	slice := array[0:3]
	fmt.Println(slice, "cap():", cap(slice), "len():", len(slice), "cap(slice)==len(array)", cap(slice) == len(array))
	//slice can be created by using make() method e.g. make(type[],length)
	//it creates a zeroed array
	x := make([]int, 5)
	fmt.Println(x)
}
