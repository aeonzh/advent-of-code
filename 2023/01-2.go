package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convert(n string) string {
	numbersMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	if _, ok := numbersMap[n]; ok {
		return numbersMap[n]
	} else {
		return n
	}
}

func main() {
	txt, err := os.ReadFile("01.txt")
	check(err)

	lines := strings.Split(string(txt), "\n")

	firstRe := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|\d`)
	lastRe := regexp.MustCompile(`.+(one|two|three|four|five|six|seven|eight|nine|\d).*?$`)
	sum := 0

	for _, line := range lines {
		// fmt.Println(line)

		firstMatches := firstRe.FindAllString(line, -1)
		// fmt.Println(firstMatches)
		first := convert(firstMatches[0])

		lastMatches := lastRe.FindAllStringSubmatch(line, -1)
		// fmt.Println(lastMatches)
		var last string
		if len(lastMatches) > 0 {
			last = convert(lastMatches[0][1])
		} else {
			last = first
		}

		// fmt.Printf("first: %s, last: %s\n", first, last)
		digit, _ := strconv.Atoi(first + last)
		// fmt.Printf("digit: %d\n", digit)
		sum += digit
	}

	fmt.Println(sum)
}
