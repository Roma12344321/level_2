package main

import (
	"reflect"
	"testing"
)

func TestFindAnagramSets(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected map[string][]string
	}{
		{
			name:  "Test with common anagrams",
			words: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name:     "Test with no anagrams",
			words:    []string{"пятак", "яблоко", "банан"},
			expected: map[string][]string{},
		},
		{
			name:  "Test with mixed case anagrams",
			words: []string{"пятак", "Пятка", "тяпка", "Листок", "слиток", "столик"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := findAnagramSets(tt.words)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, actual)
			}
		})
	}
}
