package main

import "testing"

func TestTriangle(t *testing.T)  {
	tests := []struct{a, b, c int} {
		{3,4, 5},
		{5, 12, 13},
		{12, 35, 307},
		{30000, 40000, 50000},
	}

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

func BenchmarkTriangle(b *testing.B)  {
	m, n, p := 3, 4, 5
	actual := calcTriangle(m, n)
	if p != actual {
		b.Errorf("got %d for input %d, %d; expected %d", actual, m, n, p)
	}
}
