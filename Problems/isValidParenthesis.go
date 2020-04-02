package main

import (
	"fmt"

	stack "../DS/StackArray"
)

var parenthesisMap = map[rune]rune{
	'}': '{',
	')': '(',
	']': '['}

func main() {
	var value = "()[]{}"
	var isParenthisValue1 = isValid(value)
	fmt.Println("Result of the the function: ", isParenthisValue1)
}

func isValid(s string) bool {

	var stackParenthesis = stack.New(len(s))
	for _, char := range s {
		if parenthesisMap[char] != 0 {
			var topValue rune
			if stackParenthesis.IsEmpty() {
				topValue = '#'
			} else {
				topValue = stackParenthesis.Pop().(rune)
			}

			if topValue != parenthesisMap[char] {
				return false
			}
		} else {
			stackParenthesis.Push(char)
		}
	}

	return stackParenthesis.IsEmpty()
}
