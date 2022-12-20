package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	windowLeftBorder := 0
	windowRightBorder := 0
	mp := make(map[uint8]bool)
	maxLength := 0
	for windowRightBorder < len(s) {
		if _, exist := mp[s[windowRightBorder]]; exist {
			delete(mp, s[windowLeftBorder])
			windowLeftBorder += 1
			continue
		}
		if windowRightBorder-windowLeftBorder+1 > maxLength {
			maxLength = windowRightBorder - windowLeftBorder + 1
		}
		mp[s[windowRightBorder]] = true
		windowRightBorder += 1
	}
	return maxLength
}

func main() {
	fmt.Printf("lengthOfLongestSubstring:%v", lengthOfLongestSubstring("pwwkew"))
}
