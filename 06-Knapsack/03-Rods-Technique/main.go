package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const numItems = 40

const minValue = 1
const maxValue = 10
const minWeight = 4
const maxWeight = 10
var allowedWeight int

type Item struct {
	id, blockedBy int
	blockList []int // Other items that this one blocks
	value, weight int
	isSelected bool
}

func makeItems(numItems, minValue, maxValue, minWeight, maxWeight int) []Item {
	// Initialize a pseudorandom number generator.
	// Initialize with a changing seed.
	// random := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Initialize with a fixed seed.
	random := rand.New(rand.NewSource(1337))
	
	items := make([]Item, numItems)
	for i := 0; i < numItems; i++ {
		items[i] = Item {
			id: i,
			blockedBy: -1,
			blockList: nil,
			value: random.Intn(maxValue - minValue + 1) + minValue,
			weight: random.Intn(maxWeight - minWeight + 1) + minWeight,
			isSelected: false,
		}
	}
	return items
}

// Build the items' block lists.
func makeBlockLists(items []Item) {
	for i := range items {
		items[i].blockList = []int{}
		for j := range items {
			if i != j && items[i].value >= items[j].value &&
					items[i].weight <= items[j].weight {
				items[i].blockList = append(items[i].blockList, j)
			}
		}
	}
}

// Block items on this item's blocks list.
func blockItems(source Item, items []Item) {
	for _, item := range source.blockList {
		if items[item].blockedBy == -1 {
			items[item].blockedBy = source.id
		}
	}
}

// Unblock items on this item's block list.
func unblockItems(source Item, items []Item) {
	for _, item := range source.blockList {
		if items[item].blockedBy == source.id {
			items[item].blockedBy = -1
		}
	}
}

// Print the selected items.
func printSelected(items []Item) {
	numPrinted := 0
	for i, item := range items {
		if item.isSelected {
			fmt.Printf("%d(%d, %d) ", i, item.value, item.weight)
		}
		numPrinted += 1
		if numPrinted > 100 {
			fmt.Println("...")
			return
		}
	}
	fmt.Println()
}

// Return a copy of the items slice.
func copyItems(items []Item) []Item {
	newItems := make([]Item, len(items))
	copy(newItems, items)
	return newItems
}

// Return the total value of the items.
// If addAll is false, only add up the selected items.
func sumValues(items []Item, addAll bool) int {
	total := 0
	for i := 0; i < len(items); i++ {
		if addAll || items[i].isSelected {
			total += items[i].value
		}
	}
	return total
}

// Return the total weight of the items.
// If addAll is false, only add up the selected items.
func sumWeights(items []Item, addAll bool) int {
	total := 0
	for i := 0; i < len(items); i++ {
		if addAll || items[i].isSelected {
			total += items[i].weight
		}
	}
	return total
}

// Return the value of this solution.
// If the solution is too heavy, return -1 so we prefer an empty solution.
func solutionValue(items []Item, allowedWeight int) int {
	// If the solution's total weight > allowedWeight,
	// return 0 so we won't use this solution.
	if sumWeights(items, false) > allowedWeight { return -1 }
	
	// Return the sum of the selected values.
	return sumValues(items, false)
}

// Run the algorithm. Display the elapsed time and solution.
func runAlgorithm(alg func([]Item, int) ([]Item, int, int), items []Item,
		allowedWeight int) {
	// Copy the items so the run isn't influenced by a previous run.
	testItems := copyItems(items)
	
	start := time.Now()
	
	// Run the algorithm.
	solution, totalValue, functionCalls := alg(testItems, allowedWeight)
	
	elapsed := time.Since(start)
	
	fmt.Printf("Elapsed: %f\n", elapsed.Seconds())
	printSelected(solution)
	fmt.Printf("Value: %d, Weight: %d, Calls: %d\n",
		totalValue, sumWeights(solution, false), functionCalls)
	fmt.Println()
}

// Recursively assign values in or out of the solution.
// Return the best assignment, value of that assignment,
// and the number of function calls we made.
func rodsTechnique(items []Item, allowedWeight int) ([]Item, int, int) {
	makeBlockLists(items)
	bestValue := 0
	currentValue := 0
	currentWeight := 0
	remainingValue := sumValues(items, true)
	return doRodsTechnique(items, allowedWeight,
		bestValue, currentValue, currentWeight, remainingValue, 0)
}

// Recursively assign values in or out of the solution.
// Return the best assignment, value of that assignment,
// and the number of function calls we made.
func rodsTechniqueSorted(items []Item, allowedWeight int) ([]Item, int, int) {
	makeBlockLists(items)
	// Sort the items so that the item with the longest block list
	// will come first.
	sort.Slice(items, func(i, j int) bool {
		return len(items[i].blockList) > len(items[j].blockList)
	})
	// After sorting, set the id field to the correct value.
	for i, item := range items {
		item.id = i
	}
	// Because sorting shuffled the items, make block lists again.
	makeBlockLists(items)
	bestValue := 0
	currentValue := 0
	currentWeight := 0
	remainingValue := sumValues(items, true)
	return doRodsTechnique(items, allowedWeight,
		bestValue, currentValue, currentWeight, remainingValue, 0)
}

func doRodsTechnique(items []Item, allowedWeight,
	bestValue, currentValue, currentWeight, remainingValue int,
	nextIndex int) (
		[]Item, int, int) {
	if nextIndex >= len(items) {
		testItems := copyItems(items)
		value := solutionValue(testItems, allowedWeight)
		return testItems, value, 1
	}
	if currentValue + remainingValue <= bestValue {
		return nil, 0, 1
	}
	var test1Solution []Item
	test1Solution = nil
	test1Value := 0
	test1Calls := 1
	test2Solution := copyItems(items)
	var test2Value, test2Calls int
	if currentWeight + items[nextIndex].weight <= allowedWeight &&
			items[nextIndex].blockedBy == -1 {
		items[nextIndex].isSelected = true
		test1Solution, test1Value, test1Calls =
			doRodsTechnique(items, allowedWeight,
				bestValue, currentValue + items[nextIndex].value,
				currentWeight + items[nextIndex].weight,
				remainingValue - items[nextIndex].value,
				nextIndex + 1)
	}
	
	if currentValue + remainingValue - items[nextIndex].value > bestValue {
		blockItems(items[nextIndex], items)
		items[nextIndex].isSelected = false
		test2Solution, test2Value, test2Calls =
			doRodsTechnique(items, allowedWeight,
				bestValue, currentValue,
				currentWeight,
				remainingValue - items[nextIndex].value,
				nextIndex + 1)
		if currentValue + remainingValue - items[nextIndex].value <= bestValue {
			test2Solution = nil
			test2Value = 0
		}
		unblockItems(items[nextIndex], items)
	}
	if test1Value >= test2Value {
		return test1Solution, test1Value, test1Calls + test2Calls + 1
	} else {
		return test2Solution, test2Value, test1Calls + test2Calls + 1
	}
}

func main() {
	//items := makeTestItems()
	items := makeItems(numItems, minValue, maxValue, minWeight, maxWeight)
	allowedWeight := sumWeights(items, true) / 2
	
	// Display basic parameters.
	fmt.Println("*** Parameters ***")
	fmt.Printf("# items: %d\n", numItems)
	// fmt.Printf("  all items: %v\n", items)
	fmt.Printf("Total value: %d\n", sumValues(items, true))
	fmt.Printf("Total weight: %d\n", sumWeights(items, true))
	fmt.Printf("Allowed weight: %d\n", allowedWeight)
	fmt.Println()
	
	// Exhaustive search
	if numItems > 85 { // Only run search if numItems <= 23.
		fmt.Println("Too many items Rod's technique")
	} else {
		fmt.Println("*** Rod's Technique ***")
		runAlgorithm(rodsTechnique, items, allowedWeight)
	}

	if numItems > 350 { // Only run search if numItems <= 350.
		fmt.Println("Too many items Rod's technique (Sorted speedup)")
	} else {
		fmt.Println("*** Rod's Technique (Sorted speedup) ***")
		runAlgorithm(rodsTechniqueSorted, items, allowedWeight)
	}
}

