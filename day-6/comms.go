package main

import (
	"fmt"
	"os"
)

func main() {
	dat := readInputFile()
	fmt.Println(findStartOfInterest(dat, 4))
	fmt.Println(findStartOfInterest(dat, 14))

}

func findStartOfInterest(s string, numChars int) int {
	chars := map[string]bool{}
	p0 := 0
	p1 := 0

	for p1 < len(s) {
		// check if char at p1 matches any
		// in the chars map. If it does
		// set the pointers equal to each other?
		if chars[string(s[p1])] {
			// reset
			chars[string(s[p0])] = false
			p0++
		} else {
			chars[string(s[p1])] = true
			if p1-p0 == numChars-1 {
				return p1 + 1
			}
			p1++
		}
	}

	return -1
}

func readInputFile() string {
	dat, err := os.ReadFile("day-6/comms_input.txt")
	if err != nil {
		panic(err)
	}
	return string(dat)
}
