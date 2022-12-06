package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	l1 int
	l2 int
	r1 int
	r2 int
}

func main() {
	content := readInputFile()
	pairs := contentToPairs(content)
	part1Count := 0
	part2Count := 0
	for _, p := range pairs {
		if p.hasFullOverlap() {
			part1Count++
		}

		if p.hasAnyOverlap() {
			part2Count++
		}
	}

	fmt.Println("Part 1: ", part1Count)
	fmt.Println("Part 2: ", part2Count)

}

func contentToPairs(content string) []*pair {
	lines := strings.Split(content, "\r\n")
	pairs := []*pair{}
	for _, line := range lines {
		firstDash := strings.Index(line, "-")
		comma := strings.Index(line, ",")
		lastDash := strings.LastIndex(line, "-")

		l1, err := strconv.Atoi(line[0:firstDash])
		if err != nil {
			panic(err)
		}

		l2, err := strconv.Atoi(line[firstDash+1 : comma])
		if err != nil {
			panic(err)
		}

		r1, err := strconv.Atoi(line[comma+1 : lastDash])
		if err != nil {
			panic(err)
		}

		r2, err := strconv.Atoi(line[lastDash+1:])
		if err != nil {
			panic(err)
		}

		pairs = append(pairs, &pair{
			l1: l1,
			l2: l2,
			r1: r1,
			r2: r2,
		})
	}

	return pairs
}

func (p *pair) hasFullOverlap() bool {
	return (p.l1 <= p.r1 && p.r2 <= p.l2) || (p.r1 <= p.l1 && p.l2 <= p.r2)
}

func (p *pair) hasAnyOverlap() bool {
	// if either of the points of either pair is between the points
	// of the other one

	return (p.l1 <= p.r1 && p.r1 <= p.l2) || (p.l1 <= p.r2 && p.r2 <= p.l2) || (p.r1 <= p.l1 && p.l1 <= p.r2) || (p.r1 <= p.l2 && p.l2 <= p.r2)
}

func readInputFile() string {
	dat, err := os.ReadFile("day-4/cleaning_pairs.txt")
	if err != nil {
		panic(err)
	}

	return string(dat)
}
