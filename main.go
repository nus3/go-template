package main

import (
	"fmt"
)

func main() {
	var value string
	fmt.Scan(&value)

	var count, maxLength int
	var isMatched bool

	targetChars := [4]string{"A", "T", "C", "G"}

	for _, char := range value {
		isMatched = false
		for _, targetChar := range targetChars {
			if targetChar == string(char) {
				isMatched = true
			}
		}

		if isMatched {
			count++
		} else {
			count = 0
		}

		if count > maxLength {
			maxLength = count
		}
	}

	fmt.Println(maxLength)
}
