package main

import "fmt"

func main() {

	// Creating an array
	arr := [6]string{"This", "is", "a", "Go", "interview", "question"}

	// Print array
	fmt.Println("Original Array:", arr)

	// Create a slice
	slicedArr := arr[1:4]

	// Display slice
	fmt.Println("Sliced Array:", slicedArr)

	// Length of slice calculated using len()
	fmt.Println("Length of the slice: %d", len(slicedArr))

	// Capacity of slice calculated using cap()
	fmt.Println("Capacity of the slice: %d", cap(slicedArr))
}
