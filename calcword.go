package main

import (
	"fmt"
)

func main() {
	word := "selamat malam"
	charCount := make(map[string]int)

	for i := 0; i < len(word); i++ {
		char := string(word[i])
		charCount[char]++
		fmt.Println(char)
	}

	
	fmt.Println(charCount)
}