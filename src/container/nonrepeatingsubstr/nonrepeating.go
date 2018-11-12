package main

var lastOccurred = make([]int, 0xffff)
func lengthOfNonRepeatingSubStr(s string) int  {
	for i := range lastOccurred {
		lastOccurred[i] = 0
	}
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI := lastOccurred[ch]; lastI > start {
			start = lastI
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i + 1
	}
	return maxLength
}

