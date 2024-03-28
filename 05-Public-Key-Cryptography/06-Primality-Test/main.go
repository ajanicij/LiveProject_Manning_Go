package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Initialize a pseudorandom number generator.
var random = rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed

const numTests = 20

// Return a pseudo random number in the range [min, max).
func randRange(min int, max int) int {
	// fmt.Printf("randRange: min=%d max=%d\n", min, max)
	return min + random.Intn(max - min)
}

// Perform tests to see if a number is (probably) prime.
func isProbablyPrime(p int, numTests int) bool {
	// For p < 4 we don't do the random test because randRange(3, p) would panic.
	if p < 4 {
		return p == 2 || p == 3
	}
	for i := 0; i < numTests; i++ {
		var randNum int
		for {
			// fmt.Printf("calling randRange: p=%d\n", p)
			randNum = randRange(3, p)
			// fmt.Printf("isProbablyPrime: randNum=%d\n", randNum)
			if randNum % 2 != 0 {
				break
			}
		}
		powered := fastExpMod(randNum, p - 1, p)
		if powered != 1 {
			return false
		}
	}
	return true
}

func fastExpMod(num, pow, mod int) int {
	// fmt.Printf("fastExpMod: num=%d pow=%d mod=%d\n", num, pow, mod)
	result := 1
	for pow > 0 {
		if pow % 2 == 1 {
			result = (result * num) % mod
		}
		pow /= 2
		num = (num * num) % mod
	}
	return result
}

// Probabilistically find a prime number within the range [min, max).
func findPrime(min, max, numTests int) int {
	for {
		candidate := randRange(min, max)
		if isProbablyPrime(candidate, numTests) {
			return candidate
		}
	}
}

func testKnownValues() {
	primes := []int {
		10009, 11113, 11699, 12809, 14149,
		15643, 17107, 17881, 19301, 19793,
	}
	composites := []int {
		10323, 11397, 12212, 13503, 14599,
		16113, 17547, 17549, 18893, 19999,
	}

	fmt.Println("Primes:")
	for _, p := range primes {
		if isProbablyPrime(p, numTests) {
			fmt.Printf("%d: Prime\n", p)
		} else {
			fmt.Printf("%d: Composite (?)\n", p)
		}
	}

	fmt.Println("Composites:")
	for _, p := range composites {
		if isProbablyPrime(p, numTests) {
			fmt.Printf("%d: Prime (?)\n", p)
		} else {
			fmt.Printf("%d: Composite\n", p)
		}
	}
}

func main() {
	fmt.Println("Testing for known values")
	testKnownValues()
	prob := (1.0 - math.Pow(2.0, -float64(numTests))) * 100.0
	fmt.Printf("Probability: %f%%\n", prob)
	
	for {
		var numDigits int
		fmt.Print("# Digits: ")
		fmt.Scanf("%d", &numDigits)
		if numDigits < 1 {
			break
		}
		max := int(math.Pow(10.0, float64(numDigits)))
		prime := findPrime(max/10, max, numTests)
		fmt.Printf("Prime: %d\n\n", prime)
	}
}

