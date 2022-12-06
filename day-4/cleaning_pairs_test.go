package main

import "testing"

func TestHasFullOverlap(t *testing.T) {
	testCases := []struct {
		l1   int
		r1   int
		l2   int
		r2   int
		want bool
	}{
		{
			l1:   0,
			l2:   3,
			r1:   3,
			r2:   6,
			want: false,
		},
		{
			l1:   0,
			l2:   3,
			r1:   2,
			r2:   6,
			want: false,
		},
		{
			l1:   0,
			l2:   3,
			r1:   1,
			r2:   2,
			want: true,
		},
		{
			l1:   1203,
			l2:   4434,
			r1:   3456,
			r2:   3732,
			want: true,
		},
		{
			r1:   1203,
			r2:   4434,
			l1:   3456,
			l2:   3732,
			want: true,
		},
	}

	for _, tc := range testCases {
		pair := &pair{
			l1: tc.l1,
			l2: tc.l2,
			r1: tc.r1,
			r2: tc.r2,
		}
		got := pair.hasFullOverlap()

		if tc.want != got {
			t.Errorf("hasFullOverlap(%d, %d, %d, %d) want %t got %t", tc.l1, tc.l2, tc.r1, tc.r2, tc.want, got)
		}
	}
}

func TestHasPartialOverlap(t *testing.T) {
	testCases := []struct {
		l1   int
		r1   int
		l2   int
		r2   int
		want bool
	}{
		{
			l1:   0,
			l2:   3,
			r1:   3,
			r2:   6,
			want: true,
		},
		{
			l1:   0,
			l2:   3,
			r1:   2,
			r2:   6,
			want: true,
		},
		{
			l1:   0,
			l2:   3,
			r1:   1,
			r2:   2,
			want: true,
		},
		{
			l1:   1203,
			l2:   4434,
			r1:   3456,
			r2:   3732,
			want: true,
		},
		{
			r1:   1203,
			r2:   4434,
			l1:   3456,
			l2:   3732,
			want: true,
		},
		{
			r1:   0,
			r2:   4,
			l1:   46346,
			l2:   46674,
			want: false,
		},
	}

	for _, tc := range testCases {
		pair := &pair{
			l1: tc.l1,
			l2: tc.l2,
			r1: tc.r1,
			r2: tc.r2,
		}
		got := pair.hasAnyOverlap()

		if tc.want != got {
			t.Errorf("hasAnyOverlap(%d, %d, %d, %d) want %t got %t", tc.l1, tc.l2, tc.r1, tc.r2, tc.want, got)
		}
	}
}
