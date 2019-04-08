package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLine(sc *bufio.Scanner) string {
	sc.Scan()
	return strings.TrimSpace(sc.Text())
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

	sc := bufio.NewScanner(os.Stdin)

	nmc := explodeInt(" ", readLine(sc))
	n, m, c = nmc[0], nmc[1], nmc[2]

	bValues = explodeInt(" ", readLine(sc))

	for i := 0; i < n; i++ {
		aValues := explodeInt(" ", readLine(sc))

		sum := 0
		for j := 0; j < m; j++ {
			sum += bValues[j] * aValues[j]
		}

		if sum+c > 0 {
			outputValue++
		}
	}

	fmt.Println(outputValue)
}
