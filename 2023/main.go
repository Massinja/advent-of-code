package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var filename = "dayOneInput"

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

func main() {
	fmt.Println("Hi, Elviie!")

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
			twoDigitInt := getNum(line)
			answer += twoDigitInt
		}
	}
	fmt.Println("The answer is:", answer)
}
