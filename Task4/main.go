package main

import "fmt"
// я так думаю
func printMultiplicationTable(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%d\t", i*j)
		}
		fmt.Println()
	}
}

func main() {
	number := 16
	printMultiplicationTable(number)
}
