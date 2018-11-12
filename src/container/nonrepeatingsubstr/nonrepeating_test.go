package main

import (
	"testing"
)

func TestSubStr(t *testing.T)  {
	tests := []struct{
		s string
		ans int
	} {
		{"abcabcbb", 3},
		{"pwwkew", 3},
		{"", 0},
		{"b", 1},
		{"abcabcabcd", 4},

		{"这里是慕课网", 6},
	}

	for _, tt := range tests {
		actual:= lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubStr(b *testing.B)  {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))
	ans := 8
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expecteed %d", actual, s, ans)
		}
	}
}
