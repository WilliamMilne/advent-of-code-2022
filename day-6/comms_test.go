package main

import "testing"

func TestFindStartOfPacket(t *testing.T) {
	testCases := []struct {
		in   string
		want int
	}{
		{
			in:   "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want: 7,
		},
		{
			in:   "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want: 5,
		},
		{
			in:   "nppdvjthqldpwncqszvftbrmjlhg",
			want: 6,
		},
		{
			in:   "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want: 10,
		},
		{
			in:   "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want: 11,
		},
	}

	for _, tc := range testCases {
		got := findStartOfInterest(tc.in, 4)

		if got != tc.want {
			t.Errorf("got %d want %d", got, tc.want)
		}
	}
}
