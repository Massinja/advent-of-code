package main

import "testing"

func TestParseLine(t *testing.T) {
	line := "Game 8: 9 green, 1 red; 18 green, 2 red, 7 blue; 1 blue, 9 green, 3 red; 3 red, 15 blue, 18 green"
	exp := []string{"9", "green", "1", "red", "18", "green", "2", "red", "7", "blue", "1", "blue", "9", "green", "3", "red", "3", "red", "15", "blue", "18", "green"}
	lineID := 8

	id, out := parseLine(line)
	if id != lineID {
		t.Errorf("Wrong line id")
	}
	if len(out) < len(exp) {
		t.Errorf("The set is truncated")
	}
	if len(out) == len(exp) {
		for i, v := range out {
			if exp[i] != v {
				t.Errorf("Didn't parse the line correctly: expected: %v, got: %v", exp[i], v)
			}
		}
	}
}
