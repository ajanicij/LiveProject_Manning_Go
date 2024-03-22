package main

import (
    "fmt"
)

func main() {
	for {
		var a, b int
		fmt.Print("a=")
		fmt.Scanf("%d", &a)
		if a < 1 {
			break
		}
		fmt.Print("b=")
		fmt.Scanf("%d", &b)
		if b < 1 {
			break
		}
		// fmt.Printf("a=%d, b=%d\n", a, b)
		gcd := gcd(a, b)
		lcm := lcm(a, b)
		fmt.Printf("GCD(%d, %d) = %d\n", a, b, gcd)
		fmt.Printf("LCM(%d, %d) = %d\n", a, b, lcm)
	}
}

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
