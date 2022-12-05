package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	packingList := readInputFile()
	rucksackContents := strings.Split(packingList, "\n")
	fmt.Println(rucksackContents)
	var sum int32 = 0
	for _, s := range rucksackContents {
		sum += findOverlappingItemType(s)
	}

	fmt.Printf("part 1: %d", sum)

	fmt.Printf("part 2: %d", findGroupBadgeSum(rucksackContents))
}

func findGroupBadgeSum(rucksackContents []string) int32 {
	var sum int32
	for i := 0; i < len(rucksackContents); i = i + 3 {
		var elfGroup []string = rucksackContents[i : i+3]
		sum += findBadgeOfThreeElves(elfGroup)
	}

	return sum
}

func findBadgeOfThreeElves(rucksackContents []string) int32 {
	if len(rucksackContents) != 3 {
		panic(fmt.Sprintf("rucksack contents need length of three, %d", len(rucksackContents)))
	}
	m0 := countInstancesOfItems(rucksackContents[0])
	m1 := countInstancesOfItems(rucksackContents[1])
	m2 := countInstancesOfItems(rucksackContents[2])

	for c := range m0 {
		has1 := m1[c] > 0
		has2 := m2[c] > 0
		if has1 && has2 {
			fmt.Println("=====")
			fmt.Println(c)
			fmt.Println(m0[c])
			fmt.Println(m1[c])
			fmt.Println(m2[c])
			fmt.Println("=====")
			return c
		}
	}

	return 0
}

func countInstancesOfItems(rucksackContents string) map[int32]int {
	m := make(map[int32]int)
	for _, c := range rucksackContents {
		// was somehow getting a "carriage return" in
		// the contents, so skipped that by ignoring 13.
		// not an ideal long-term solution, but did solve the problem appropriately.
		if c == 13 {
			continue
		}
		converted := convertASCIIValueToPriorityValue(c)
		fmt.Println(converted)
		m[converted] = m[converted] + 1
	}
	return m
}

func convertASCIIValueToPriorityValue(c int32) int32 {
	if c >= 97 && c <= 122 {
		return c - 96
	}

	if c >= 65 && c <= 90 {
		return c - 38
	}

	panic(c)
}

func findOverlappingItemType(s string) int32 {
	compartmentSize := len(s) / 2
	items := make(map[int32]bool)

	compartment1 := s[0:compartmentSize]
	for _, c := range compartment1 {
		items[c] = true
	}

	compartment2 := s[compartmentSize:]
	for _, c := range compartment2 {
		if items[c] {
			return convertASCIIValueToPriorityValue(c)
		}
	}

	return -1
}

func readInputFile() string {
	dat, err := os.ReadFile("day-3/packing_list.txt")
	if err != nil {
		panic(err)
	}

	return string(dat)
}
