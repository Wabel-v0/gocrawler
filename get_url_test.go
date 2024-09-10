package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {

	tests := []struct {
		name      string
		inputUrl  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputUrl: "https://baba.json.me",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>baba.json</span>
				</a>
				<a href="https://other.com/path/one">
					<span>baba.json</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://baba.json.me/path/one", "https://other.com/path/one"},
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := getURLsFromHTML(tc.inputBody, tc.inputUrl)
			want := tc.expected
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Test %d: got %q, want %q", i, got, want)
			}

		})
	}

}
