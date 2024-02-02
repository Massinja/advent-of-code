package main

import "testing"

func TestGetNum(t *testing.T) {
	lines := []string{"gldsixrhss186seven6", "gnpksz4", "4919", "pbc19", "gjjkdsg"}
	answers := []int{16, 44, 49, 19, 0}
	for x := 0; x < len(lines); x++ {
		number := getNum([]byte(lines[x]))
		if number != answers[x] {
			t.Errorf("Line: %s, expected: %d, got: %d", lines[x], answers[x], number)
		}
	}

}

func TestGetNum2(t *testing.T) {
	lines := []string{"aaon7one718onegfqtdbtxfcmd", "threeninedtr7219", "two2geight", "1b9four", "123four5six", "123four5"}
	answers := []int{71, 39, 28, 14, 16, 15}
	for x := 0; x < len(lines); x++ {
		number := getNum2([]byte(lines[x]))
		if number != answers[x] {
			t.Errorf("Line: %s, expected: %d, got: %d", lines[x], answers[x], number)
		}
	}
}
