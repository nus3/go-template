package main

import "fmt"

func main() {
	var value string
	fmt.Scan(&value)

	basePattern := map[string]string{
		"A": "T",
		"T": "A",
		"G": "C",
		"C": "G",
	}

	fmt.Println(basePattern[value])
}
