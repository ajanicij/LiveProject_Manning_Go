package main

import (
	"fmt"
	"strconv"
	"math/rand"
)

// Perform linear search.
// Return the target's location in the slice and the number of tests.
// If the item is not found, return -1 and the number tests.
func linearSearch(slice []int, target int) (index, numTests int) {
	numTests = 0
	for index := 0; index < len(slice); index++ {
		numTests++
		if slice[index] == target {
			return index, numTests
		}
	}
	return -1, numTests
}

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
	fmt.Printf("%v\n", slice[:numItems])
}

func main() {
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)

	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()
	
	for {
		var target_str string
		fmt.Printf("target: ")
		fmt.Scanln(&target_str)
		if target_str == "" {
			fmt.Println("Exit")
			break
		}
		if target, err := strconv.Atoi(target_str); err == nil {
			index, numTests := linearSearch(slice, target)
			if index == -1 {
				fmt.Printf("Target %d not found, %d tests\n",
					target, numTests)
			} else {
				fmt.Printf("values[%d] = %d, %d tests\n",
					index, target, numTests)
			}
		} else {
			fmt.Printf("Error: %s\n", err)
		}
	}
}

