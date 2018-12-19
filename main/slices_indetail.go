package main

import "fmt"

func main() {
	//this chapter explains slices in detail
	//make a nil slice
	mySlice := make([]int, 0)
	fmt.Println(mySlice)
	//append a single value to the slice
	myAppendedSlice := append(mySlice, 1)
	//access the newly added element
	fmt.Println(myAppendedSlice[0])
	//the underlying array of the slice is adjusted as new elements are added
	myAppendedSlice1 := append(mySlice, 3, 2, 1)
	fmt.Println(myAppendedSlice1)
	//make can work on multi dimension arrays
	//the below example shows an slice having [3]int elements
	tictactoe := make([][3]int, 2)
	fmt.Println(tictactoe)
	tictactoe[0][0] = 1
	fmt.Println(tictactoe)
	//append one element of type [3]int to the slice
	completeTicTacToe := append(tictactoe, [3]int{0, 1, 0})
	fmt.Println(completeTicTacToe)
	//we can create a zeroed & nil slice of slice
	sliceOfSlice := make([][]int, 0)
	fmt.Println(sliceOfSlice)
	//append to the inner slice
	sliceOfSliceWithOneElement := append(sliceOfSlice, []int{1, 2, 3})
	fmt.Println(sliceOfSliceWithOneElement)
}
