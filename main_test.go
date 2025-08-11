package main

import "testing"

type CleanTestCase struct {
	input    string
	expected string
}

func TestCleanInput(t *testing.T) {
	cases := []CleanTestCase{
		{
			input:    "1 + 2   + 3  ",
			expected: "1+2+3",
		},
		{
			input:    "  3 /4 + 3 * 21",
			expected: "3/4+3*21",
		},
		{
			input:    " 25 *23 + 0 / 1   ",
			expected: "25*23+0/1",
		},
		{
			input:    " 2514  **2 %2   ",
			expected: "2514**2%2",
		},
		{
			input:    "( ( 23 * 3 ) + 21   * 3 ) %2",
			expected: "((23*3)+21*3)%2",
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("length mismatch between %v and %v! got: %d, want: %d\n", actual, c.expected, len(actual), len(c.expected))
		}

	}
}

type ParseTestCase struct {
	raw       string
	formatted []string
}

func TestParseExpression(t *testing.T) {
	// ...
}
