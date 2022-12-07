// 1657. Determine if Two Strings Are Close
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(closeStrings("abc", "bca"))
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	var alphabet1 [26]bool
	var alphabet2 [26]bool
	wc1, wc2 := make(map[int32]int), make(map[int32]int)

	for _, v := range word1 {
		alphabet1[v-'a'] = true
		wc1[v-'a']++
	}

	for _, v := range word2 {
		alphabet2[v-'a'] = true
		wc2[v-'a']++
	}

	for i := 0; i < 26; i++ {
		if alphabet1[i] != alphabet2[i] {
			return false
		}
	}

	if len(wc1) != len(wc2) {
		return false
	}

	var r1, r2 []int
	for _, v := range wc1 {
		r1 = append(r1, v)
	}
	for _, v := range wc2 {
		r2 = append(r2, v)
	}
	sort.Ints(r1)
	sort.Ints(r2)

	for i := 0; i < len(r1); i++ {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}
