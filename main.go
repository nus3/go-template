package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func readInt() (int, error) {
	inputValue := readLine()
	return strconv.Atoi(strings.TrimSpace(inputValue))
}

func explodeString(delimiter string, inputValue string) []string {
	splittedValue := strings.Split(inputValue, delimiter)

	var trimmedValues []string

	for _, value := range splittedValue {
		if value != "" {
			trimmedValues = append(trimmedValues, value)
		}
	}

	return trimmedValues
}

func explodeInt(delimiter string, inputValue string) []int {
	inputStrings := explodeString(" ", inputValue)

	var splittedInts []int

	for _, inputString := range inputStrings {
		var (
			intValue int
			err      error
		)
		intValue, err = strconv.Atoi(inputString)

		if err != nil {
			panic(err)
		}
		splittedInts = append(splittedInts, intValue)
	}

	return splittedInts
}

func main() {
	var (
		n, m, c, outputValue int
		bValues              []int
	)

	nmc := explodeInt(" ", readLine())

	n, m, c = nmc[0], nmc[1], nmc[2]

	bValues = explodeInt(" ", readLine())

	for i := 0; i < n; i++ {
		aValues := explodeInt(" ", readLine())

		sum := c
		for mIndex := 0; mIndex < m; mIndex++ {
			sum += bValues[mIndex] * aValues[mIndex]
		}

		if sum > 0 {
			outputValue++
		}
	}

	fmt.Println(outputValue)
}
