// 451. Sort Characters By Frequency
package main

import (
	"strings"
)

func main() {
	frequencySort("ettree")
}

func frequencySort(s string) string {
	m := make(map[int32]int)
	arr := []string{}
	str := ""
	for _, v := range s {
		m[v]++
	}

	for key, val := range m {
		s := ""
		for i := 0; i < val; i++ {
			s += string(key)
		}
		arr = append(arr, s)
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if len(arr[i]) < len(arr[j]) {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	str = strings.Join(arr, "")

	return str
}
