package main

import (
 "fmt"
 "math"
)
func gcd(a, b int) int {
	if b == 0 {
	 return a
	}
	return gcd(b, a%b)
   }
   
   func findNumbers(num int) []int {
	divisors := []int{}
	for i := 1; i <= int(math.Sqrt(float64(num))); i++ {
	 if num%i == 0 {
	  divisors = append(divisors, i)
	  if i != num/i {
	   divisors = append(divisors, num/i)
	  }
	 }
	}
	return divisors
   }
   
   func findCommonDivisors(arr []int) []int {
   
	var commonDivisors []int
	result := arr[0]
	for _, num := range arr[1:] {
	 result = gcd(result, num)
	}
   
	divisors := findNumbers(result)
   
	for _, divisor := range divisors {
	 allDivisible := true
	 for _, num := range arr {
	  if num%divisor != 0 {
	   allDivisible = false
	   break
	  }
	 }
	 if allDivisible {
	  commonDivisors = append(commonDivisors, divisor)
	 }
	}
   
	return commonDivisors
   }
   
   func main() {
	numbers := []int{42, 12, 18}
	fmt.Println(findCommonDivisors(numbers))
   }
