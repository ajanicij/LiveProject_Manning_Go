package main

import (
	"fmt"
	"os"
)

// Build a sieve of Eratosthenes.
func sieveOfEratosthenes(max int) []bool {
	isPrime := make([]bool, max + 1)
	for i := 2; i <= max; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= max; i++ {
		if !isPrime[i] {
			continue
		}
		for j := i*i; j <= max; j = j + i {
			isPrime[j] = false
		}
	}
	return isPrime
}

// Build an Euler's sieve.
func eulersSieve(max int) []bool {
	isPrime := make([]bool, max + 1)
	for i := 2; i <= max; i++ {
		isPrime[i] = true
	}
	for p := 3; p*p <= max; p += 2 {
		if !isPrime[p] {
			continue
		}
		maxQ := max / p
		if maxQ % 2 == 0 {
			maxQ -= 1
		}
		for q := maxQ; q >= p; q-- {
			if isPrime[q] {
				isPrime[q*p] = false
			}
		}
	}
	return isPrime
}

// Print out the primes in the sieve.
func printSieve(sieve []bool) {
	count := 1
	fmt.Printf("%d ", 2)
	for i := 3; i < len(sieve); i += 2 {
		if sieve[i] {
			fmt.Printf("%d ", i)
			count++
			if count == 10 {
				count = 0
				fmt.Println()
			}
		}
	}
	fmt.Println()
}

func main() {
	var max int
	fmt.Print("max: ")
	fmt.Scanf("%d", &max)
	if max < 2 {
		os.Exit(0)
	}
	sieve := sieveOfEratosthenes(max)
	sieveEuler := eulersSieve(max)
	fmt.Printf("Result of Sieve of Eratosthenes for max=%d:\n", max)
	printSieve(sieve)
	
	fmt.Printf("Result of Euler's Sieve for max=%d:\n", max)
	printSieve(sieveEuler)
}
