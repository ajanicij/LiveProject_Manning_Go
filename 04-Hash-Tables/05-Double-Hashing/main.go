package main

import (
	"fmt"
	"math/rand"
	"time"
)

// djb2 hash function. See http://www.cse.yorku.ca/~oz/hash.html.
func hash1(value string) int {
	hash := 5381
	for _, ch := range value {
		hash = ((hash << 5) + hash) + int(ch)
	}

	// Make sure the result is non-negative.
	if hash < 0 {
		hash = -hash
	}
	return hash
}

// Jenkins one_at_a_time hash function.
// See https://en.wikipedia.org/wiki/Jenkins_hash_function
func hash2(value string) int {
    hash := 0
    for _, ch := range value {
        hash += int(ch);
        hash += hash << 10
        hash ^= hash >> 6
    }

    // Make sure the result is non-negative.
    if hash < 0 { hash = -hash }

    // Make sure the result is not 0.
    if hash == 0 { hash = 1 }
    return hash;
}

type Employee struct {
	name    string
	phone   string
	deleted bool
}

type DoubleHashTable struct {
	capacity  int
	employees []*Employee
}

// Initialize a DoubleHashTable and return a pointer to it.
func NewDoubleHashTable(capacity int) *DoubleHashTable {
	table := &DoubleHashTable{
		capacity: capacity,
	}

	employees := make([]*Employee, capacity)
	table.employees = employees
	return table
}

// Display the hash table's contents.
func (hashTable *DoubleHashTable) dump() {
	for i, employee := range hashTable.employees {
		if employee != nil {
			if employee.deleted {
				fmt.Printf("%d: xxx\n", i)
			} else {
				fmt.Printf("%d: %s\t%s\n", i, employee.name, employee.phone)
			}
		} else {
			fmt.Printf("%d: ---\n", i)
		}
	}
}

// Return the key's index or where it would be if present and
// the probe sequence length.
// If the key is not present and the table is full, return -1 for the index.
func (hashTable *DoubleHashTable) find(name string) (int, int) {
	hash1 := hash1(name) % hashTable.capacity
	hash2 := hash2(name) % hashTable.capacity
	deletedIndex := -1
	for i := 0; i < hashTable.capacity; i++ {
		index := (hash1 + i*hash2) % hashTable.capacity // Double hashing
		employee := hashTable.employees[index]
		if employee == nil {
			if deletedIndex >= 0 {
				return deletedIndex, i + 1
			} else {
				return index, i + 1
			}

		}
		if employee.deleted {
			if deletedIndex == -1 {
				deletedIndex = index
			}
		} else {
			if employee.name == name {
				return index, i + 1
			}
		}
	}
	if deletedIndex >= 0 {
		return deletedIndex, hashTable.capacity
	}
	return -1, hashTable.capacity
}

// Show this key's probe sequence.
func (hashTable *DoubleHashTable) probe(name string) int {
	// Hash the key.
	hash1 := hash1(name) % hashTable.capacity
	hash2 := hash2(name) % hashTable.capacity
	fmt.Printf("Probing %s (%d, %d)\n", name, hash1, hash2)

	// Keep track of a deleted spot if we find one.
	deletedIndex := -1

	// Probe up to hashTable.capacity times.
	for i := 0; i < hashTable.capacity; i++ {
		index := (hash1 + i*hash2) % hashTable.capacity

		fmt.Printf("    %d: ", index)
		if hashTable.employees[index] == nil {
			fmt.Printf("---\n")
		} else if hashTable.employees[index].deleted {
			fmt.Printf("xxx\n")
		} else {
			fmt.Printf("%s\n", hashTable.employees[index].name)
		}

		// If this spot is empty, the value isn't in the table.
		if hashTable.employees[index] == nil {
			// If we found a deleted spot, return its index.
			if deletedIndex >= 0 {
				fmt.Printf("    Returning deleted index %d\n", deletedIndex)
				return deletedIndex
			}

			// Return this index, which holds nil.
			fmt.Printf("    Returning nil index %d\n", index)
			return index
		}

		// If this spot is deleted, remember where it is.
		if hashTable.employees[index].deleted {
			if deletedIndex < 0 {
				deletedIndex = index
			}
		} else if hashTable.employees[index].name == name {
			// If this cell holds the key, return its data.
			fmt.Printf("    Returning found index %d\n", index)
			return index
		}

		// Otherwise continue the loop.
	}

	// If we get here, then the key is not
	// in the table and the table is full.

	// If we found a deleted spot, return it.
	if deletedIndex >= 0 {
		fmt.Printf("    Returning deleted index %d\n", deletedIndex)
		return deletedIndex
	}

	// There's nowhere to put a new entry.
	fmt.Printf("    Table is full\n")
	return -1
}

// Add an item to the hash table.
func (hashTable *DoubleHashTable) set(name string, phone string) {
	index, _ := hashTable.find(name)
	if index < 0 {
		panic("table full")
	}
	employee := hashTable.employees[index]
	if employee == nil || employee.deleted {
		employee = &Employee{
			name:  name,
			phone: phone,
		}
		hashTable.employees[index] = employee
	} else {
		employee.phone = phone
	}
}

// Return an item from the hash table.
func (hashTable *DoubleHashTable) get(name string) string {
	index, _ := hashTable.find(name)
	if index < 0 {
		return ""
	}
	employee := hashTable.employees[index]
	if employee == nil || employee.deleted {
		return ""
	}
	return employee.phone
}

// Return true if the person is in the hash table.
func (hashTable *DoubleHashTable) contains(name string) bool {
	index, _ := hashTable.find(name)
	if index < 0 {
		return false
	}
	employee := hashTable.employees[index]
	if employee == nil || employee.deleted {
		return false
	}
	return true
}

func (hashTable *DoubleHashTable) delete(name string) {
	index, _ := hashTable.find(name)
	if index >= 0 {
		employee := hashTable.employees[index]
		if employee != nil {
			employee.deleted = true
		}
	}
}

// Make a display showing whether each array entry is nil.
func (hashTable *DoubleHashTable) dumpConcise() {
	// Loop through the array.
	for i, employee := range hashTable.employees {
		if employee == nil {
			// This spot is empty.
			fmt.Printf(".")
		} else if employee.deleted {
			// Deleted entry.
			fmt.Printf("X")
		} else {
			// Display this entry.
			fmt.Printf("O")
		}
		if i%50 == 49 {
			fmt.Println()
		}
	}
	fmt.Println()
}

// Return the average probe sequence length for the items in the table.
func (hashTable *DoubleHashTable) aveProbeSequenceLength() float32 {
	totalLength := 0
	numValues := 0
	for _, employee := range hashTable.employees {
		if employee != nil {
			_, probeLength := hashTable.find(employee.name)
			totalLength += probeLength
			numValues++
		}
	}
	return float32(totalLength) / float32(numValues)
}

func main() {
	// Make some names.
	employees := []Employee{
		Employee{"Ann Archer", "202-555-0101", false},
		Employee{"Bob Baker", "202-555-0102", false},
		Employee{"Cindy Cant", "202-555-0103", false},
		Employee{"Dan Deever", "202-555-0104", false},
		Employee{"Edwina Eager", "202-555-0105", false},
		Employee{"Fred Franklin", "202-555-0106", false},
		Employee{"Gina Gable", "202-555-0107", false},
	}

	hashTable := NewDoubleHashTable(10)
	for _, employee := range employees {
		hashTable.set(employee.name, employee.phone)
	}
	hashTable.dump()

	hashTable.probe("Hank Hardy")
	fmt.Printf("Table contains Sally Owens: %t\n", hashTable.contains("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", hashTable.contains("Dan Deever"))
	fmt.Println("Deleting Dan Deever")
	hashTable.delete("Dan Deever")
	fmt.Printf("Table contains Dan Deever: %t\n", hashTable.contains("Dan Deever"))
	fmt.Printf("Sally Owens: %s\n", hashTable.get("Sally Owens"))
	fmt.Printf("Fred Franklin: %s\n", hashTable.get("Fred Franklin"))
	fmt.Println("Changing Fred Franklin")
	hashTable.set("Fred Franklin", "202-555-0100")
	fmt.Printf("Fred Franklin: %s\n", hashTable.get("Fred Franklin"))
	hashTable.dump()

	hashTable.probe("Ann Archer")
	hashTable.probe("Bob Baker")
	hashTable.probe("Cindy Cant")
	hashTable.probe("Dan Deever")
	hashTable.probe("Edwina Eager")
	hashTable.probe("Fred Franklin")
	hashTable.probe("Gina Gable")
	hashTable.set("Hank Hardy", "202-555-0108")
	hashTable.probe("Hank Hardy")

	// Look at clustering.
	fmt.Println(time.Now())                   // Print the time so it will compile if we use a fixed seed.
	random := rand.New(rand.NewSource(12345)) // Initialize with a fixed seed
	// random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed
	bigCapacity := 1009
	bigHashTable := NewDoubleHashTable(bigCapacity)
	numItems := int(float32(bigCapacity) * 0.9)
	for i := 0; i < numItems; i++ {
		str := fmt.Sprintf("%d-%d", i, random.Intn(1000000))
		bigHashTable.set(str, str)
	}
	bigHashTable.dumpConcise()
	fmt.Printf("Average probe sequence length: %f\n",
		bigHashTable.aveProbeSequenceLength())
}
