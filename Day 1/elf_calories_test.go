package main

import "testing"

func TestElfCalories(t *testing.T) {
	testCases := []struct {
		calories string
		want     int
	}{
		{
			calories: "123\n456\n\n1000",
			want:     1000,
		},
	}

	for _, tc := range testCases {
		got, err := findElfWithMostCalories(tc.calories)
		if err != nil {
			t.Fatalf("findElfWithMostCalories got unexpected error")
		}
		if tc.want != got {
			t.Errorf("findElfWithMostCalories() want %d got %d", tc.want, got)
		}
	}
}

func TestThreeElvesCalories(t *testing.T) {
	testCases := []struct {
		calories string
		want     int
	}{
		{
			calories: "600\n400\n\n1000\n\n400",
			want:     2400,
		},
		{
			calories: "1\n\n10\n20\n30\n40\n\n67\n33\n\n4000",
			want:     4200,
		},
	}

	for _, tc := range testCases {
		got, err := findTotalOfThreeElvesWithMostCalories(tc.calories)
		if err != nil {
			t.Fatalf("findElfWithMostCalories got unexpected error")
		}
		if tc.want != got {
			t.Errorf("findElfWithMostCalories() want %d got %d", tc.want, got)
		}

	}
}
