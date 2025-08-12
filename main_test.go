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
	sectioned []string
}

func TestParseExpression(t *testing.T) {

	cases := []ParseTestCase{
		{
			raw:       "1+2-3",
			sectioned: []string{"1", "+", "2", "-", "3"},
		},
		{
			raw:       "3/4-3*21",
			sectioned: []string{"3", "/", "4", "-", "3", "*", "21"},
		},
		{
			raw:       "25*23+0/1",
			sectioned: []string{"25", "*", "23", "+", "0", "/", "1"},
		},
		{
			raw:       "2514**2%2",
			sectioned: []string{"2514", "**", "2", "%", "2"},
		},
	}

	for _, c := range cases {
		parsed := parseExpression(c.raw)

		if len(parsed) != len(c.sectioned) {
			t.Errorf("length mismatch between %v and %v! got: %d, want: %d\n", parsed, c.sectioned, len(parsed), len(c.sectioned))
		}

		for i, elem := range parsed {
			if elem != c.sectioned[i] {
				t.Errorf("section mismatch! got: %v, want: %v\nslice: %v\n", elem, c.sectioned[i], parsed)
			}
		}
	}

}

// TODO: Come up with a better way to parse. Think: binary expression tree
type ParseTestCaseN struct {
	raw       string
	sectioned map[string][]string
}
