// 1704. Determine if String Halves Are Alike
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(halvesAreAlike("textbook"))
}

func halvesAreAlike(s string) bool {
	var vowels = "aeiouAEIOU"
	var leftCount, rightCount int
	for i := 0; i < len(s)/2; i++ {
		//Left
		if ok := strings.Contains(vowels, string(s[i])); ok {
			leftCount++
		}
		//Right
		if ok := strings.Contains(vowels, string(s[len(s)-1-i])); ok {
			rightCount++
		}
	}
	return leftCount == rightCount
}
