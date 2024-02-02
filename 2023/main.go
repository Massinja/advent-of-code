package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var filename = "dayOneInput"

var digits = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// func getNum takes a slice of bytes and returns a two digit number
// it makes out of the first and last digits found in the line
func getNum(line []byte) int {
	lineLength := len(line)
	var digitOne, digitTwo uint8
	twoDigitInt := 0
	// get the first digit
	for x := 0; x < lineLength; x++ {
		if unicode.IsDigit(rune(line[x])) {
			digitOne = line[x]
			x = lineLength
		}
	}
	// get the last digit
	for x := lineLength - 1; x >= 0; x-- {
		if unicode.IsDigit(rune(line[x])) {
			digitTwo = line[x]
			x = -1
		}
	}
	// if at least a single digit was found, make a two digit number
	if digitOne != 0 {
		twoDigitStr := string(digitOne) + string(digitTwo)
		twoDigitInt, _ = strconv.Atoi(twoDigitStr)
	}
	return twoDigitInt
}

// func getNum2 takes a slice of bytes and returns a two digit number
// it makes out of the first and last digits found in the line
// considers strings like "one", "two", etc - as a valid digit
func getNum2(line []byte) int {

	var digitOne, digitTwo, twoDigitInt int
	digitOne = -1

	for len(line) > 0 {
		for word, num := range digits {

			a := []byte(strconv.Itoa(num))[0]
			if line[0] == a {
				if digitOne == -1 {
					digitOne = num
				}
				digitTwo = num
			} else if len(line) >= len(word) {
				if string(line[:len(word)]) == word {
					if digitOne == -1 {
						digitOne = num
					}
					digitTwo = num
				}
			}
		}
		line = line[1:]
	}
	if digitOne != -1 {
		twoDigitInt = digitOne*10 + digitTwo
	}
	return twoDigitInt
}

func main() {
	fmt.Println("Hi, Elviie!")

	second := flag.Bool("second", false, "solve second part of the challenge")
	flag.Parse()

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Couldn't open file %s: %v\n", filename, err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	answer := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) > 0 {
			if !*second {
				answer += getNum(line)
			} else {
				answer += getNum2(line)
			}
		}
	}
	fmt.Println("The answer is:", answer)
}
