package main

import (
	"reflect"
	"testing"
)

func TestPrintReport(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []Page
	}{
		{
			name: "sorts pages by number of visits",
			input: map[string]int{
				"https://baba.json.me":  1,
				"https://baba.json.com": 2,
				"https://baba.json.org": 3,
			},
			expected: []Page{
				{Key: "https://baba.json.org", Value: 3},
				{Key: "https://baba.json.com", Value: 2},
				{Key: "https://baba.json.me", Value: 1},
			},
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := sortedPages(tc.input)
			want := tc.expected
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Test %d: got %q, want %q", i, got, want)
			}
		})
	}

}
