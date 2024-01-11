package main

import (
	"sort"
	"strings"
)

func rearrangeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	counter := make(map[rune]int)
	for _, c := range s {
		counter[c]++
	}
	sortable := make([]rune, 0, len(counter))
	for c := range counter {
		sortable = append(sortable, c)
	}
	sort.Slice(sortable, func(i, j int) bool {
		return counter[sortable[i]] > counter[sortable[j]]
	})
	res := make([]rune, 0, len(s))
	indexes := make([]int, len(sortable))
	for i := range sortable {
		indexes[i] = counter[sortable[i]]
	}
	for {
		i := -1
		for j, count := range indexes {
			if count > 0 {
				i = j
				break
			}
		}
		if i == -1 {
			break
		}
		res = append(res, sortable[i])
		indexes[i]--
		for i := range indexes {
			if i == 0 || indexes[i-1] == 0 {
				continue
			}
			indexes[i-1] = max(indexes[i-1]-1, 0)
		}
	}
	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	test := []string{"aab", "aaab", "aaaaabbb"}
	for _, t := range test {
		res := rearrangeString(t)
		println(t, "->", res)
	}
}