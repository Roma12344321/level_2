package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func findAnagramSets(words []string) map[string][]string {
	anagramMap := make(map[string][]string)
	for _, word := range words {
		lower := strings.ToLower(word)
		sorted := sortString(lower)
		anagramMap[sorted] = append(anagramMap[sorted], lower)
	}
	res := make(map[string][]string)
	for _, v := range anagramMap {
		if len(v) > 1 {
			for _, word := range words {
				lower := strings.ToLower(word)
				if slices.Contains(v, lower) {
					sort.Strings(v)
					res[lower] = v
					break
				}
			}
		}
	}
	return res
}

func sortString(s string) string {
	arr := []rune(s)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return string(arr)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagramSets := findAnagramSets(words)
	for key, set := range anagramSets {
		fmt.Printf("Key: %s, Set: %v\n", key, set)
	}
}
