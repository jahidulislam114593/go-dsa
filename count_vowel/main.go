package main

import "fmt"

func countVowels(s string) int {
	cnt := 0
	for _, ch := range s {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			cnt++
		}

	}

	return cnt
}
func main() {
	str := "Jahidul islam"
	cnt := countVowels(str)
	fmt.Println(cnt)
}
