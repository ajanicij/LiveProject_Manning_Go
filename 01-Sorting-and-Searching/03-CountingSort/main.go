package main

import (
	"fmt"
	"math/rand"
)

type Customer struct {
	id           string
	numPurchases int
}

func makeRandomSlice(numItems, max int) []Customer {
	s := make([]Customer, numItems)
	for i := range s {
		s[i] = Customer{
			id:           fmt.Sprintf("C%d", i),
			numPurchases: rand.Intn(max),
		}
	}
	return s
}

// Print at most numItems items.
func printSlice(slice []Customer, numItems int) {
	if len(slice) < numItems {
		numItems = len(slice)
	}
	for _, x := range slice[:numItems] {
		fmt.Printf("%v", x)
	}
}

// Verify that the slice is sorted.
func checkSorted(slice []Customer) {
	numItems := len(slice) - 1
	for i := range slice[:numItems] {
		if slice[i].numPurchases > slice[i+1].numPurchases {
			fmt.Println("Not sorted")
			return
		}
	}
	fmt.Println("Sorted")
	return
}

func countingSort(slice []Customer, max int) []Customer {
	counts := make([]int, max+1)

	// Loop through the items and for each item increment
	// the corresponding count.
	for _, item := range slice {
		index := item.numPurchases
		counts[index]++
	}

	// Convert the counts int to the number of items less than
	// or equal to each value.
	for i := 0; i <= max; i++ {
		if i > 0 {
			counts[i] += counts[i-1]
		}
	}

	// Generate sorted slice.
	sliceB := make([]Customer, len(slice))
	for i := len(slice) - 1; i >= 0; i-- {
		key := slice[i].numPurchases
		index := counts[key] - 1
		sliceB[index] = slice[i]
		counts[key]--
	}
	return sliceB
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	sorted := countingSort(slice, max)
	printSlice(sorted, 40)

	// Verify that it's sorted.
	checkSorted(sorted)
}
