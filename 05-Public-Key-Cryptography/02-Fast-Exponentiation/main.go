package main

import (
	"fmt"
	"math"
)

func fastExp(num, pow int) int {
	result := 1
	for pow > 0 {
		if pow % 2 == 1 {
			result *= num
		}
		pow /= 2
		num *= num
	}
	return result
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

func main() {
	for {
		var num, pow, mod int
		fmt.Print("num: ")
		fmt.Scanf("%d", &num)
		if num < 1 {
			break
		}

		fmt.Print("pow: ")
		fmt.Scanf("%d", &pow)
		if pow < 1 {
			break
		}

		fmt.Print("mod: ")
		fmt.Scanf("%d", &mod)
		if mod < 1 {
			break
		}

		res := fastExp(num, pow)
		fmt.Printf("%d ^ %d = %d\n", num, pow, res)
		check := int(math.Pow(float64(num), float64(pow)))
		fmt.Printf("  -- check: %d\n", check)

		res = fastExpMod(num, pow, mod)
		fmt.Printf("%d ^ %d (mod %d) = %d\n", num, pow, mod, res)
		check = check % mod
		fmt.Printf("  -- check: %d\n", check)
	}
}
