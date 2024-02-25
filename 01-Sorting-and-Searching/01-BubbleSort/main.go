package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Make a slice containing pseudorandom numbers in [0, max).
func makeRandomSlice(numItems, max int) []int {
	s := make([]int, numItems)
	for i := range s {
		s[i] = rand.Intn(max)
	}
	return s
}

// Print at most numItems items.
func printSlice(slice []int, numItems int) {
	if len(slice) < numItems {
		numItems = len(slice)
	}
	for _, x := range(slice[:numItems]) {
		fmt.Println(x)
	}
}

// Verify that the slice is sorted.
func checkSorted(slice []int) {
	numItems := len(slice) - 1
	for i := range(slice[:numItems]) {
		if slice[i] > slice[i+1] {
			fmt.Println("Not sorted")
			return
		}
	}
	fmt.Println("Sorted")
	return
}

func bubbleSort(slice []int) {
	n := len(slice)
	for {
		swapped := false
		for i := range(slice[:n-1]) {
			if slice[i] > slice[i+1] {
				temp := slice[i]
				slice[i] = slice[i+1]
				slice[i+1] = temp
				swapped = true
			}
		}
		if !swapped {
			break
		}
		n -= 1
	}
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display an unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	start := time.Now()
	bubbleSort(slice)
	t := time.Now()
	printSlice(slice, 40)
	fmt.Printf("Time: %v\n", t.Sub(start))

	// Verify that it's sorted.
	checkSorted(slice)
}
