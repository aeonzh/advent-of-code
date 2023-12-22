package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	txt, err := os.ReadFile("01.txt")
	check(err)

	firstDigitsRe := regexp.MustCompile(`(?m:^\D*(\d))`)
	firstDigitSlice := firstDigitsRe.FindAllSubmatch([]byte(txt), -1)

	lastDigitsRe := regexp.MustCompile(`(?m:(\d)\D*$)`)
	lastDigitSlice := lastDigitsRe.FindAllSubmatch([]byte(txt), -1)

	var lines int
	if len(firstDigitSlice) == len(lastDigitSlice) {
		lines = len(firstDigitSlice)
	}

	var sum int = 0
	for i := 0; i < lines; i++ {
		digit, _ := strconv.Atoi(string(firstDigitSlice[i][1]) + string(lastDigitSlice[i][1]))
		sum += digit
	}

	fmt.Println(sum)
}
