package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var filename = "day2input"
var cubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// returns ID number of the processed line, and a slice of strings.
// each string is either a number of cubes or their corresponding colour
// input line is in the following format:
// "Game 8: 9 green, 1 red; 18 green, 2 red, 7 blue; 1 blue, 9 green, 3 red; 3 red, 15 blue, 18 green"
func parseLine(line string) (int, []string) {
	s := strings.Split(line, ": ")
	index := strings.Split(s[0], " ")
	lineID, _ := strconv.Atoi(index[1])
	re := regexp.MustCompile("[;:, ]+")
	split := re.Split(s[1], -1)
	set := []string{}
	for i := range split {
		set = append(set, split[i])
	}
	return lineID, set
}

// checks if number of cubes of a certain color is less than the number in var cubes
// returns id of the line if passes check
// returns 0 if doesn't pass
func validateSet(lineID int, set []string) int {
loop:
	for x := 0; x < len(set); {
		colour := set[x+1]
		num, _ := strconv.Atoi(set[x])
		if num > cubes[colour] {
			lineID = 0
			break loop
		}
		x += 2
	}
	return lineID
}

// func getMinNumOfCubes returns a minimum number of cubes
// that are necessary to play the game
func getMinNumOfCubes(set []string) map[string]int {
	cubes := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for x := 0; x < len(set); {
		colour := set[x+1]
		num, _ := strconv.Atoi(set[x])
		if num > cubes[colour] {
			cubes[colour] = num
		}
		x += 2
	}
	return cubes
}

// func multCubeNums returns the result of multiplication
// of all the map values
func multCubeNums(cubes map[string]int) int {
	num := 1
	for _, v := range cubes {
		num = num * v
	}
	return num
}

func main() {
	fmt.Println("Hi, Elviie!")
	idTotal := 0
	powerSum := 0
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Couldn't open file %s: %v\n", filename, err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		lineID, set := parseLine(line)
		lineID = validateSet(lineID, set)
		idTotal += lineID
		cubes := getMinNumOfCubes(set)
		gamePower := multCubeNums(cubes)
		powerSum += gamePower
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Couldn't read properly from %s: %v\n", filename, err)
	}

	fmt.Println("Sum of eligible IDs is", idTotal)
	fmt.Println("Sum of the power of the game sets is", powerSum)
}
