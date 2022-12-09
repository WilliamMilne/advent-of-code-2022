package main

import (
	"testing"
)

func makeStack(strings []string) Stack {
	var stack Stack
	for _, s := range strings {
		stack.Push(s)
	}
	return stack
}

func TestParsingCommand(t *testing.T) {
	testCases := []struct {
		command string
		want    *CommandData
	}{
		{
			command: "move 1 from 2 to 1",
			want: &CommandData{
				Source: 2,
				Dest:   1,
				Qty:    1,
			},
		},
	}

	for _, tc := range testCases {
		got := parseCommand(tc.command)

		if got.Dest != tc.want.Dest {
			t.Errorf("Dest mismatch")
		}

		if got.Source != tc.want.Source {
			t.Errorf("Source mismatch")
		}

		if got.Qty != tc.want.Qty {
			t.Errorf("Qty mismatch")
		}
	}
}

func TestParsingStacks(t *testing.T) {
	testCases := []struct {
		stackLines []string
		wantStacks []Stack
	}{
		{
			stackLines: []string{
				"[A] [D] [B]",
				" 1   2   3 ",
			},
			wantStacks: []Stack{
				makeStack([]string{"A"}),
				makeStack([]string{"D"}),
				makeStack([]string{"B"}),
			},
		},
		{
			stackLines: []string{
				"    [J]    ",
				"    [E] [F]",
				"[A] [D] [B]",
				" 1   2   3 ",
			},
			wantStacks: []Stack{
				makeStack([]string{"A"}),
				makeStack([]string{"D", "E", "J"}),
				makeStack([]string{"B", "F"}),
			},
		},
	}

	for _, tc := range testCases {
		got := parseStackLines(tc.stackLines)
		if len(got) == 0 {
			t.Fatalf("no stacks returned")
		}

		for i, stack := range tc.wantStacks {
			gotStack := got[i]

			for !stack.IsEmpty() {
				wantEl, _ := stack.Pop()
				gotEl, _ := gotStack.Pop()

				if gotEl != wantEl {
					t.Errorf("wanted %q got %q", wantEl, gotEl)
				}
			}
		}
	}
}
