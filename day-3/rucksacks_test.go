package main

import "testing"

func TestRucksackSorting(t *testing.T) {
	testCases := []struct {
		s    string
		want int32
	}{
		{
			s:    "abcade",
			want: 1,
		},
	}

	for _, tc := range testCases {
		got := findOverlappingItemType(tc.s)

		if tc.want != got {
			t.Errorf("findOverlappingItemType(%q) want %d got %d", tc.s, tc.want, got)
		}
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		s    []string
		want int32
	}{
		{
			s:    []string{"aA", "aB", "aC"},
			want: 1,
		},
		{
			s:    []string{"bA", "bB", "bC"},
			want: 2,
		},
		{
			s:    []string{"bA", "aA", "cA"},
			want: 27,
		},
		{
			s:    []string{"aaaaB", "aaaaC", "aaaaA"},
			want: 1,
		},
		{
			s: []string{"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg"},
			want: 18,
		},
	}

	for _, tc := range testCases {
		got := findBadgeOfThreeElves(tc.s)
		t.Errorf("hi")
		if tc.want != got {
			t.Errorf("findBadgeOfThreeElves(%q) want %d got %d", tc.s, tc.want, got)
		}
	}
}

func TestFindGroupBadgeSum(t *testing.T) {
	testCases := []struct {
		s    []string
		want int32
	}{
		{
			s: []string{"vJrwpWtwJgWrhcsFMMfFFhFp\n",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\n",
				"PmmdzqPrVvPwwTWBwg\n"},
			want: 18,
		},
		{
			s: []string{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\n",
				"ttgJtRGJQctTZtZT\n",
				"CrZsJsPPZsGzwwsLwLmpwMDw\n"},
			want: 52,
		},
		{
			s: []string{
				"vJrwpWtwJgWrhcsFMMfFFhFp\n",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\n",
				"PmmdzqPrVvPwwTWBwg\n",
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\n",
				"ttgJtRGJQctTZtZT\n",
				"CrZsJsPPZsGzwwsLwLmpwMDw\n",
			},
			want: 70,
		},
	}

	for _, tc := range testCases {
		got := findGroupBadgeSum(tc.s)
		if tc.want != got {
			t.Errorf("findGroupBadgeSum(%q) want %d got %d", tc.s, tc.want, got)
		}
	}
}
