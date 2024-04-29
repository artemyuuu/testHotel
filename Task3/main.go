package main

import (
	"fmt"
)

func findPrimesInRange(minNum, maxNum int) []int {
	primes := []int{}
	for num := minNum; num <= maxNum; num++ {
		if isPrime(num) {
			primes = append(primes, num)
		}
	}
	return primes
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	i := 5
//решето Эратосфена
	for i*i <= num {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func main() {
	minNum := 11
	maxNum := 12341212312312
	primes := findPrimesInRange(minNum, maxNum)
	fmt.Printf("Простые числа в диапазоне от %d до %d: %v\n", minNum, maxNum, primes)
}
