package main

import "testing"

type TestCase struct {
	input    string
	expected []string
}

func TestCleanInput(t *testing.T) {
	cases := []TestCase{
		{
			input:    "1 + 2   + 3  ",
			expected: []string{"1", "+", "2", "+", "3"},
		},
		{
			input:    "  3 /4 + 3 * 21",
			expected: []string{"3", "/", "4", "3", "*", "21"},
		},
		{
			input:    " 25 *23 + 0 / 1   ",
			expected: []string{"25", "*", "23", "+", "0", "/", "1"},
		},
		{
			input:    " 2514  **2 %2   ",
			expected: []string{"2514", "**", "%", "2"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("length mismatch between %v and %v! got: %d, want: %d\n", actual, c.expected, len(actual), len(c.expected))
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("symbol mismatch between %s and %s!\n", actual[i], c.expected[i])
			}
		}

	}
}
