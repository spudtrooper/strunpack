package strunpack

import (
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
)

func TestUnpackerFromRegex(t *testing.T) {
	type Typ struct {
		Name string
		Age  int
	}
	tests := []struct {
		input    string
		re       *regexp.Regexp
		expected Typ
		err      error
	}{
		{
			input: "Alice 25",
			re:    regexp.MustCompile(`(?P<Name>\w+) (?P<Age>\d+)`),
			expected: Typ{
				Name: "Alice",
				Age:  25,
			},
		},
		{
			input: "Alice 25",
			re:    regexp.MustCompile(`(\w+) (\d+)`),
			expected: Typ{
				Name: "Alice",
				Age:  25,
			},
		},
	}

	for _, test := range tests {
		unpacker := FromRegex[Typ](test.re)
		result, err := unpacker.Unpack(test.input)

		if !errors.Is(err, test.err) {
			t.Errorf("Input: %s\nExpected error: %v\nGot error: %v", test.input, test.err, err)
		}

		if !cmp.Equal(*result, test.expected) {
			t.Errorf("Input: %s\nExpected: %+v\nGot: %+v", test.input, test.expected, result)
		}
	}
}

func TestUnpackerFromString(t *testing.T) {
	type Typ struct {
		Name string
		Age  int
	}
	tests := []struct {
		input    string
		re       string
		expected Typ
		err      error
	}{
		{
			input: "Alice 25",
			re:    `(?P<Name>\w+) (?P<Age>\d+)`,
			expected: Typ{
				Name: "Alice",
				Age:  25,
			},
		},
		{
			input: "Alice 25",
			re:    `(\w+) (\d+)`,
			expected: Typ{
				Name: "Alice",
				Age:  25,
			},
		},
	}

	for _, test := range tests {
		unpacker := FromString[Typ](test.re)
		result, err := unpacker.Unpack(test.input)

		if !errors.Is(err, test.err) {
			t.Errorf("Input: %s\nExpected error: %v\nGot error: %v", test.input, test.err, err)
		}

		if !cmp.Equal(*result, test.expected) {
			t.Errorf("Input: %s\nExpected: %+v\nGot: %+v", test.input, test.expected, result)
		}
	}
}
