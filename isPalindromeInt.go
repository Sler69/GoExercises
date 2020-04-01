package main

import "fmt"

func main() {
	var lol = 1222222221
	var result = isPalindrome(lol)
	fmt.Println("result", result)
}

func isPalindrome(value int) bool {

	if value < 0 || (value%10 == 0 && value != 0) {
		return false
	}

	var lastHalf = 0
	for value > lastHalf {
		lastHalf = (lastHalf * 10) + value%10
		value = value / 10
	}
	return false
}
