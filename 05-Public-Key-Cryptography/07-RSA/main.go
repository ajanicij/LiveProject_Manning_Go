package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Initialize a pseudorandom number generator.
var random = rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for {
		if a == 0 {
			return b
		}
		if b == 0 {
			return a
		}
		rem := a % b
		a = b
		b = rem
	}
}

func lcm(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	gcd := gcd(a, b)
	return a * (b / gcd)
}

// Calculate the totient function λ(n)
// where n = p * q and p and q are prime.
func totient(p, q int) int {
	l := lcm(p - 1, q - 1)
	return l
}

// Return a pseudo random number in the range [min, max).
func randRange(min int, max int) int {
	return min + random.Intn(max - min)
}

// Pick a random exponent e in the range (2, λn)
// such that gcd(e, λn) = 1.
func randomExponent(λn int) int {
	for {
		e := randRange(3, λn)
		if gcd(e, λn) == 1 {
			return e
		}
	}
}

// From Wikipedia
// https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm#Computing_multiplicative_inverses_in_modular_structures
func inverseMod_old(a, n int) int {
	fmt.Printf("inverseMod: a=%d n=%d\n", a, n)

	t := 0
	newt := 1
	r := n
	newr := a
	
	fmt.Printf("inverseMod: t=%d newt=%d r=%d newr=%d\n",
		t, newt, r, newr)
	
	for newr != 0 {
		quotient := r / newr
		t, newt = newt, t - quotient * newt
		r, newr = newr, r - quotient * newr
		fmt.Printf("inverseMod (loop): quotient=%d, t=%d newt=%d r=%d newr=%d\n",
			quotient, t, newt, r, newr)
	}
	
	if r > 1 {
		return -1
	}
	if t < 0 {
		t += n
	}
	
	return t
}

// Slightly modified from inverseMod_old for clarity.
func inverseMod(a, n int) int {
	// fmt.Printf("inverseMod: n=%d a=%d\n", n, a)

	a1 := 1
	a2 := 0
	b1 := 0
	b2 := 1
	m1 := n
	m2 := a
	// assert a1*n + b1*a == m1
	// assert a2*n + b2*a == m2
	
	// NOTE: We don't really need a1, a2 and a3 for this algorithm.
	//       We are just calculating them for the asserts.
	//       At the point after the loop, a1*n+b1*a==1, so b1 will
	//       give us the inverse of a (provided that gcd(n, a)==1, i.e.
	//       m1==1).
	
	for m2 != 0 {
		// assert:
		//   a1*n + b1*a == m1
		//   a2*n + b2*a == m2
		q := m1 / m2
		m3 := m1 - q * m2
		a3 := a1 - q * a2
		b3 := b1 - q * b2
		
		// Because m1==a1*n+b1*a, m2==a2*n+b2*a, m3==m1-q*m2
		//         => m3==(a1*n+b1*a - q*(a2*n+b2*a)) == (a1 - q*a2)*n + (b1 - q*b2)*a ==
		//              == a3*n + b3*a
		
		// assert:
		//   a1*n + b1*a == m1
		//   a2*n + b2*a == m2
		//   a3*n + b3*a == m3
		
		m1, m2 = m2, m3
		a1, a2 = a2, a3
		b1, b2 = b2, b3
		
		// fmt.Printf("  a1=%d b1=%d m1=%d a2=%d b2=%d m2=%d\n",
		//	a1, b1, m1, a2, b2, m2)
	}
	
	if m1 > 1 { // n and a are not coprimes!
		return -1
	}
	if b1 < 0 {
		b1 += n
	}
	
	return b1
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

func main() {
	p := findPrime(10000, 50000, 40)
	q := findPrime(10000, 50000, 40)
	n := p * q
	lambdan := totient(p, q)
	var e int
	for {
		e = randRange(3, lambdan)
		if gcd(e, lambdan) == 1 {
			break
		}
	}
	d := inverseMod(e, lambdan)
	fmt.Println("*** Public ***")
	fmt.Printf("Public key modulus: %d\n", n)
	fmt.Printf("Public key exponent e: %d\n", e)
	fmt.Println()
	fmt.Println("*** Private ***")
	fmt.Printf("Primes: %d, %d\n", p, q)
	fmt.Printf("λ(n): %d\n", lambdan)
	fmt.Printf("d: %d\n", d)
	
	for {
		var m int
		fmt.Print("\nMessage: ")
		fmt.Scanf("%d", &m)
		if m < 1 {
			break
		}
		c := fastExpMod(m, e, n)
		fmt.Printf("Ciphertext: %d\n", c)
		
		d := fastExpMod(c, d, n)
		fmt.Printf("Decrypted plaintext: %d\n", d)
	}
}
