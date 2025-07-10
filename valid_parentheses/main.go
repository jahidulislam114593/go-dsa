package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}

	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]

			if (top == '(' && ch == ')') ||
				(top == '{' && ch == '}') ||
				(top == '[' && ch == ']') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	tests := []string{
		"()[]{}", // true
		"([{}])", // true
		"((]",    // false
		"({[)]}", // false
		"((()))", // true
		"({}",    // false
	}

	for _, t := range tests {
		fmt.Printf("isValid(%q) = %v\n", t, isValid(t))
	}
}
