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

func explode(delimiter string, inputValue string) []string {
	splittedValue := strings.Split(inputValue, delimiter)

	var trimmedValues []string

	for _, value := range splittedValue {
		if value != "" {
			trimmedValues = append(trimmedValues, value)
		}
	}

	return trimmedValues
}

func explodeString(delimiter string) []string {
	inputValue := readLine()
	return explode(delimiter, inputValue)
}

func explodeInt(delimiter string) []int {
	inputStrings := explodeString(" ")

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
	s := readLine()
	i, err := readInt()
	sa := explodeString(" ")
	ia := explodeInt(" ")

	fmt.Println(s)
	fmt.Println(i, err)
	fmt.Println(sa)
	fmt.Println(ia)
}
