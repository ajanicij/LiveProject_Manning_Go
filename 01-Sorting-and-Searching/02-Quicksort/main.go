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

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	i := lo - 1
	var j int
	for j = lo; j <= hi-1; j++ {
		if arr[j] <= pivot {
			i++
			swap(arr, i, j)
		}
	}
	i++
	swap(arr, i, hi)
	return i
}

func quicksort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	lo := 0
	hi := len(arr) - 1

	p := partition(arr, lo, hi)
	quicksort(arr[:p])
	quicksort(arr[p+1:])
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
	start := time.Now()
	quicksort(slice)
	t := time.Now()
	printSlice(slice, 40)
	fmt.Printf("Time: %v\n", t.Sub(start))

	// Verify that it's sorted.
	checkSorted(slice)
}
