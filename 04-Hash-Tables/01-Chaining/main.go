package main

import (
	"fmt"
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

type ChainingHashTable struct {
    numBuckets int
    buckets     [][]*Employee
}

// Initialize a ChainingHashTable and return a pointer to it.
func NewChainingHashTable(numBuckets int) *ChainingHashTable {
	buckets := make([][]*Employee, numBuckets)
	table := &ChainingHashTable{
		numBuckets: numBuckets,
		buckets: buckets,
	}
	return table
}

// Display the hash table's contents.
func (hashTable *ChainingHashTable) dump() {
	for i, bucket := range hashTable.buckets {
		fmt.Printf("Bucket %d:\n", i)
		for j := range bucket {
			employee := bucket[j]
			fmt.Printf("\t%s: %s\n", employee.name, employee.phone)
		}
	}
}

// Find the bucket and Employee holding this key.
// Return the bucket number and Employee number in the bucket.
// If the key is not present, return the bucket number and -1.
func (hashTable *ChainingHashTable) find(name string) (int, int) {
	index := hash(name) % hashTable.numBuckets
	bucket := hashTable.buckets[index]
	for i, employee := range bucket {
		if employee.name == name {
			return index, i
		}
	}
	return index, -1
}

// Add an item to the hash table.
func (hashTable *ChainingHashTable) set(name string, phone string) {
	index, employeeIndex := hashTable.find(name)
	if employeeIndex >= 0 {
		hashTable.buckets[index][employeeIndex].phone = phone
		return
	}
	// assert employeeIndex == -1
	employee := &Employee{
		name: name,
		phone: phone,
	}
	hashTable.buckets[index] = append(hashTable.buckets[index], employee)
}

// Return an item from the hash table.
func (hashTable *ChainingHashTable) get(name string) (string) {
	index, employeeIndex := hashTable.find(name)
	if employeeIndex >= 0 {
		return hashTable.buckets[index][employeeIndex].phone
	}
	// assert employeeIndex == -1
	return ""
}

// Return true if the person is in the hash table.
func (hashTable *ChainingHashTable) contains(name string) (bool) {
	_, employeeIndex := hashTable.find(name)
	return employeeIndex != -1
}

// Delete this key's entry.
func (hashTable *ChainingHashTable) delete(name string) {
	index, employeeIndex := hashTable.find(name)
	if employeeIndex >= 0 {
		hashTable.buckets[index] = append(hashTable.buckets[index][:employeeIndex],
			hashTable.buckets[index][employeeIndex+1:]...)
	}
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
        Employee { "Herb Henshaw",  "202-555-0108" },
        Employee { "Ida Iverson",   "202-555-0109" },
        Employee { "Jeb Jacobs",    "202-555-0110" },
    }
    
    hashTable := NewChainingHashTable(10)
    for _, employee := range employees {
        hashTable.set(employee.name, employee.phone)
    }
    hashTable.dump()

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
}

