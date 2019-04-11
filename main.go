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
	var total float64
	for i := 0; i < len(arr); i++ {
		total += float64(arr[i])
	}
	fmt.Println(total / float64(len(arr)))
	// using range to loop through arr
	total = 0
	for _, val := range arr { // _ notifies that the key will not be needed
		total += float64(val)
	}
	fmt.Println(total / float64(len(arr)))
}
