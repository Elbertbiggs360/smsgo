package main

import "fmt"

func main() {
	name := "Bruce"
	var lastName string = "Bigirwenkya"
	const gender string = "male"
	fmt.Println("Hello, my name is", name, lastName)
	// arrays
	var arr [5]int // are a fixed size
	arr[4] = 500   // their values must match the specified value data type
	fmt.Println(arr)
}
