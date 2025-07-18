package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello World",
			expected: []string{
				"hello", "world",
			},
		},
		{
			input: "chakrak tiyang",
			expected: []string{
				"chakrak", "tiyang",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The Lengths are not equal: %v vs %v",
				len(actual),
				len(cs.expected),
			)
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]

			if actualWord != expectedWord {
				t.Errorf("%v != %v",
					actualWord,
					expectedWord,
				)
			}
		}
	}
}
