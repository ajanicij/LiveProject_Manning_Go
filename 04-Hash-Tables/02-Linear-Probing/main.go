package main

import (
	"fmt"
	"time"
	"math/rand"
)

// djb2 hash function. See http://www.cse.yorku.ca/~oz/hash.html.
func hash(value string) int {
    hash := 5381
    for _, ch := range value {
        hash = ((hash << 5) + hash) + int(ch)
    }

    // Make sure the result is non-negative.
    if hash < 0 { hash = -hash }
    return hash
}

type Employee struct {
    name        string
    phone       string
}

type LinearProbingHashTable struct {
    capacity    int
    employees   []*Employee
}

// Initialize a LinearProbingHashTable and return a pointer to it.
func NewLinearProbingHashTable(capacity int) *LinearProbingHashTable {
	table := &LinearProbingHashTable{
		capacity: capacity,
	}
	
	employees := make([]*Employee, capacity)
	table.employees = employees
	return table
}

// Display the hash table's contents.
func (hashTable *LinearProbingHashTable) dump() {
	for i, employee := range(hashTable.employees) {
		if employee != nil {
			fmt.Printf("%d: %s\t%s\n", i, employee.name, employee.phone)
		} else {
			fmt.Printf("%d: ---\n", i)		
		}
	}
}

// Return the key's index or where it would be if present and
// the probe sequence length.
// If the key is not present and the table is full, return -1 for the index.
func (hashTable *LinearProbingHashTable) find(name string) (int, int) {
	baseIndex := hash(name) % hashTable.capacity
	for i := 0; i < hashTable.capacity; i++ {
		index := (baseIndex + i) % hashTable.capacity
		employee := hashTable.employees[index]
		if employee == nil {
			return index, i + 1
		}
		if employee.name == name {
			return index, i + 1
		}
	}
	return -1, hashTable.capacity
}

// Add an item to the hash table.
func (hashTable *LinearProbingHashTable) set(name string, phone string) {
	index, _ := hashTable.find(name)
	if index < 0 {
		panic("table full")
	}
	employee := hashTable.employees[index]
	if employee != nil {
		employee.phone = phone
	} else {
		employee = &Employee{
			name: name,
			phone: phone,
		}
		hashTable.employees[index] = employee
	}
}

// Return an item from the hash table.
func (hashTable *LinearProbingHashTable) get(name string) (string) {
	index, _ := hashTable.find(name)
	if index < 0 {
		return ""
	}
	employee := hashTable.employees[index]
	if employee == nil {
		return ""
	}
	return employee.phone
}

// Return true if the person is in the hash table.
func (hashTable *LinearProbingHashTable) contains(name string) (bool) {
	index, _ := hashTable.find(name)
	if index < 0 {
		return false
	}
	employee := hashTable.employees[index]
	return employee != nil
}

// Make a display showing whether each array entry is nil.
func (hashTable *LinearProbingHashTable) dumpConcise() {
    // Loop through the array.
    for i, employee := range hashTable.employees {
        if employee == nil {
            // This spot is empty.
            fmt.Printf(".")
        } else {
            // Display this entry.
            fmt.Printf("O")
        }
        if i % 50 == 49 { fmt.Println() }
    }
    fmt.Println()
}

// Return the average probe sequence length for the items in the table.
func (hashTable *LinearProbingHashTable) aveProbeSequenceLength() float32 {
    totalLength := 0
    numValues := 0
    for _, employee := range(hashTable.employees) {
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
    employees := []Employee {
        Employee { "Ann Archer",    "202-555-0101" },
        Employee { "Bob Baker",     "202-555-0102" },
        Employee { "Cindy Cant",    "202-555-0103" },
        Employee { "Dan Deever",    "202-555-0104" },
        Employee { "Edwina Eager",  "202-555-0105" },
        Employee { "Fred Franklin", "202-555-0106" },
        Employee { "Gina Gable",    "202-555-0107" },
    }
    
    hashTable := NewLinearProbingHashTable(10)
    for _, employee := range employees {
        hashTable.set(employee.name, employee.phone)
    }
    hashTable.dump()

    fmt.Printf("Table contains Sally Owens: %t\n", hashTable.contains("Sally Owens"))
    fmt.Printf("Table contains Dan Deever: %t\n", hashTable.contains("Dan Deever"))
    // fmt.Println("Deleting Dan Deever")
    // hashTable.delete("Dan Deever")
    // fmt.Printf("Table contains Dan Deever: %t\n", hashTable.contains("Dan Deever"))
    fmt.Printf("Sally Owens: %s\n", hashTable.get("Sally Owens"))
    fmt.Printf("Fred Franklin: %s\n", hashTable.get("Fred Franklin"))
    fmt.Println("Changing Fred Franklin")
    hashTable.set("Fred Franklin", "202-555-0100")
    fmt.Printf("Fred Franklin: %s\n", hashTable.get("Fred Franklin"))

    // Look at clustering.
    fmt.Println(time.Now()) // Print the time so it will compile if we use a fixed seed.
    random := rand.New(rand.NewSource(12345)) // Initialize with a fixed seed
    // random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed
    bigCapacity := 1009
    bigHashTable := NewLinearProbingHashTable(bigCapacity)
    numItems := int(float32(bigCapacity) * 0.9)
    for i := 0; i < numItems; i++ {
        str := fmt.Sprintf("%d-%d", i, random.Intn(1000000))
        bigHashTable.set(str, str)
    }
    bigHashTable.dumpConcise()
    fmt.Printf("Average probe sequence length: %f\n",
        bigHashTable.aveProbeSequenceLength())
}

