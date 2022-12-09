package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CommandData struct {
	Source int
	Dest   int
	Qty    int
}

// stack implementation from: https://www.educative.io/answers/how-to-implement-a-stack-in-golang
type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", fmt.Errorf("Stack is empty")
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, nil
}

func parseStackLines(lines []string) []*Stack {
	// understand how many stacks there will be
	length := len(lines)
	lastLine := lines[length-1]
	stacks := []*Stack{}
	for i, char := range lastLine {
		s := string(char)
		if s != " " {
			stack := traverseStack(lines, i)
			stacks = append(stacks, &stack)
		}
	}

	return stacks
}

func traverseStack(lines []string, c int) Stack {
	var stack Stack
	for i := len(lines) - 2; i >= 0; i-- {
		if string(lines[i][c]) == " " {
			return stack
		}
		stack.Push(string(lines[i][c]))
	}
	return stack
}

func parseCommand(s string) *CommandData {
	s = strings.Replace(s, "move ", "", -1)
	s = strings.Replace(s, "from ", "", -1)
	s = strings.Replace(s, "to ", "", -1)
	numbers := strings.Split(s, " ")
	qty, err := strconv.Atoi(numbers[0])
	if err != nil {
		panic(err)
	}
	source, err := strconv.Atoi(numbers[1])
	if err != nil {
		panic(err)
	}
	dest, err := strconv.Atoi(numbers[2])
	if err != nil {
		panic(err)
	}
	return &CommandData{
		Qty:    qty,
		Source: source,
		Dest:   dest,
	}
}

func parseCommands(lines []string) []*CommandData {
	commands := []*CommandData{}
	for _, line := range lines {
		commands = append(commands, parseCommand(line))
	}
	return commands
}

func main() {
	fileContent := readInputFile()
	lines := strings.Split(fileContent, "\r\n")
	stackInput, commandInput := splitInput(lines)
	stacks := parseStackLines(stackInput)
	commands := parseCommands(commandInput)

	simulate(stacks, commands)

	printTopOfStacks(stacks)
}

func simulate(stacks []*Stack, commands []*CommandData) {
	for _, command := range commands {
		source := command.Source - 1
		dest := command.Dest - 1
		qty := command.Qty

		// Part 1
		// for i := 0; i < qty; i++ {
		// 	if !stacks[source].IsEmpty() {
		// 		moving, _ := stacks[source].Pop()
		// 		stacks[dest].Push(moving)
		// 	}
		// }

		// Part 2
		// maybe not the most elegant, but
		// put the boxes onto a temporary stack,
		// then pop them off it to preserve order
		var tempStack Stack
		for i := 0; i < qty; i++ {
			if !stacks[source].IsEmpty() {
				moving, _ := stacks[source].Pop()
				tempStack.Push(moving)
			}
		}

		for !tempStack.IsEmpty() {
			moving, _ := tempStack.Pop()
			stacks[dest].Push(moving)
		}
	}
}

func printTopOfStacks(stacks []*Stack) {
	fmt.Println("number of stacks: " + strconv.Itoa(len(stacks)))
	for _, stack := range stacks {
		if !stack.IsEmpty() {
			result, _ := stack.Pop()
			fmt.Print(result + " ")
		} else {
			fmt.Print("Empty stack ")
		}
	}
	fmt.Println()
}

func readInputFile() string {
	dat, err := os.ReadFile("day-5/stack_input.txt")
	if err != nil {
		panic(err)
	}

	return string(dat)
}

func splitInput(lines []string) ([]string, []string) {
	for i, line := range lines {
		if line == "" {
			return lines[:i], lines[i+1:]
		}
	}

	panic("improper data format")
}
