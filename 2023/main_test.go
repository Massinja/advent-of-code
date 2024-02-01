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
