package main

import (
	"fmt"
	"os"
	"strings"
)

var part1ResultMap map[string]int
var part2ResultMap map[string]int

func initMaps() {
	part1ResultMap = make(map[string]int)
	part1ResultMap["A X"] = (3 + 1)
	part1ResultMap["A Y"] = (6 + 2)
	part1ResultMap["A Z"] = (0 + 3)

	part1ResultMap["B X"] = (0 + 1)
	part1ResultMap["B Y"] = (3 + 2)
	part1ResultMap["B Z"] = (6 + 3)

	part1ResultMap["C X"] = (6 + 1)
	part1ResultMap["C Y"] = (0 + 2)
	part1ResultMap["C Z"] = (3 + 3)

	// X - lose
	// Y - draw
	// Z - win
	part2ResultMap = make(map[string]int)
	part2ResultMap["A X"] = (0 + 3)
	part2ResultMap["A Y"] = (3 + 1)
	part2ResultMap["A Z"] = (6 + 2)

	part2ResultMap["B X"] = (0 + 1)
	part2ResultMap["B Y"] = (3 + 2)
	part2ResultMap["B Z"] = (6 + 3)

	part2ResultMap["C X"] = (0 + 2)
	part2ResultMap["C Y"] = (3 + 3)
	part2ResultMap["C Z"] = (6 + 1)
}

func main() {
	initMaps()
	strategyGuide := readInputFile()
	rounds := strings.Split(strategyGuide, "\n")

	totalScore := computeTotalScore(rounds, part1ResultMap)
	fmt.Printf("part 1: %d", totalScore)

	totalScore = computeTotalScore(rounds, part2ResultMap)
	fmt.Printf(" part 2: %d", totalScore)
}

func computeTotalScore(rounds []string, roundResults map[string]int) int {
	// Opponent choice:
	// A - Rock
	// B - Paper
	// C - Scissors

	// My choice:
	// X - Rock
	// Y - Paper
	// Z - Scissors

	// Scores per choice:
	// Rock - 1
	// Paper - 2
	// Scissors - 3

	// Scores per round result:
	// Win - 6
	// Draw - 3
	// Lose - 0
	score := 0
	for _, round := range rounds {
		round = strings.TrimSpace(round)
		roundResult := roundResults[round]
		score += roundResult
	}

	return score
}

func readInputFile() string {
	dat, err := os.ReadFile("strategy_guide.txt")
	if err != nil {
		panic(err)
	}

	return string(dat)
}
