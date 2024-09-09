package main

import (
	"testing"
)

func TestNormalizeUrl(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://baba.json.me/path",
			expected: "baba.json.me/path",
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := NormalizeUrl(tc.inputURL)
			want := tc.expected
			if got != want {
				t.Errorf("Test %d: got %q, want %q", i, got, want)

			}
		})
	}

}
