package main

import "testing"

func TestComputeTotalScore(t *testing.T) {
	testCases := []struct {
		rounds []string
		want   int
	}{
		{
			rounds: []string{
				"A Y",
				"B X",
				"C Z",
			},
			want: 15,
		},
		{
			rounds: []string{
				"A X",
				"B Y",
				"C Z",
			},
			want: 15,
		},
		{
			rounds: []string{
				"A X",
				"A X",
				"A X",
			},
			want: 12,
		},
		{
			rounds: []string{
				"A X",
			},
			want: 4,
		},
	}

	initMaps()

	for _, tc := range testCases {
		got := computeTotalScore(tc.rounds, part1ResultMap)
		if tc.want != got {
			t.Errorf("computeTotalScore(%v) want %d got %d", tc.rounds, tc.want, got)
		}
	}
}
