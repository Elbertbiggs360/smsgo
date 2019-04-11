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
	fmt.Println("for loop, data type conversion", total/float64(len(arr)))
	// using range to loop through arr
	total = 0
	for _, val := range arr { // _ notifies that the key will not be needed
		total += float64(val)
	}
	fmt.Println("range looping", total/float64(len(arr)))
	// shorter syntax for declaring array
	x := [2]float64{
		23,
		46, //trailing comma needed
	}
	fmt.Println("shorter syntax", x)
	var _ []float64                  // slice. Doesnt have a length specified
	slice1 := make([]float64, 5)     // use make syntax to create slices
	slice2 := make([]float64, 5, 10) // 10 is the underlying array size. 5 is the slice
	fmt.Println("Slices 1:", slice1, "and 2:", slice2)
	slice3 := arr[0:5]
	fmt.Println("Normal slice made from array", slice3)
}
