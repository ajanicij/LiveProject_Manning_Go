package main

import (
	"fmt"
	"time"
)

var primes []int

const maxPrimes = 2000000

func init() {
	isPrime := eulersSieve(maxPrimes)
	for i, flag := range *isPrime {
		if flag {
			primes = append(primes, i)
		}
	}
}

// Build an Euler's sieve.
func eulersSieve(max int) *[]bool {
	isPrime := make([]bool, max + 1)
	for i := 2; i <= max; i++ {
		isPrime[i] = true
	}
	for i := 4; i <= max; i += 2 {
		isPrime[i] = false
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
	return &isPrime
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

func divides(p, q int) bool {
	return q % p == 0
}

func findFactors(num int) []int {
	result := make([]int, 0)
	for divides(2, num) {
		result = append(result, 2)
		num /= 2
	}
	factor := 3
	for factor*factor <= num {
		if divides(factor, num) {
			result = append(result, factor)
			num = num / factor
		} else {
			factor += 2
		}
	}
	if num > 1 {
		result = append(result, num)
	}
	
	return result
}

func findFactorsSieve(num int) []int {
	result := make([]int, 0)
	for divides(2, num) {
		result = append(result, 2)
		num /= 2
	}
	for _, factor := range primes {
		for divides(factor, num) {
			result = append(result, factor)
			num = num / factor
		}
	}
	if num > 1 {
		result = append(result, num)
	}
	
	return result
}

func multiplySlice(slice []int) int {
	result := 1
	for _, el := range slice {
		result *= el
	}
	return result
}

func main() {
	for {
		fmt.Print("Enter num: ")
		var num int
		fmt.Scanf("%d", &num)
		if num < 2 {
			break
		}

		// Find the factors the slow way.
		start := time.Now()
		factors := findFactors(num)
		elapsed := time.Since(start)
		fmt.Printf("findFactors:       %f seconds\n", elapsed.Seconds())
		fmt.Println(multiplySlice(factors))
		fmt.Println(factors)
		fmt.Println()

		// Use the Euler's sieve to find the factors.
		start = time.Now()
		factors = findFactorsSieve(num)
		elapsed = time.Since(start)
		fmt.Printf("findFactorsSieve: %f seconds\n", elapsed.Seconds())
		fmt.Println(multiplySlice(factors))
		fmt.Println(factors)
		fmt.Println()
	}
}

