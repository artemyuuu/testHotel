package main

import (
	"fmt"
)

func pluralizeComputers(number int) string {
	lastDigit := number % 10
	lastTwoDigits := number % 100
	if lastTwoDigits != 11 && lastDigit%10 == 1 {
		return fmt.Sprintf("%d компьютер", number)
	} else if !(12 <= lastTwoDigits && lastTwoDigits <= 14) && lastDigit >= 2 && lastDigit <= 4 {
		return fmt.Sprintf("%d компьютера", number)
	} else {
		return fmt.Sprintf("%d компьютеров", number)
	}
}



func main() {
	fmt.Println(pluralizeComputers(1)) 
	fmt.Println(pluralizeComputers(3)) 
	fmt.Println(pluralizeComputers(4))
	fmt.Println(pluralizeComputers(14))
	fmt.Println(pluralizeComputers(12))
	fmt.Println(pluralizeComputers(5))
	fmt.Println(pluralizeComputers(19))
	fmt.Println(pluralizeComputers(11))
	fmt.Println(pluralizeComputers(21))
	fmt.Println(pluralizeComputers(10229))
}
