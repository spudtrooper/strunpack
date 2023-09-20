package strunpack

import (
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
)

func TestFromString(t *testing.T) {
	tests := []struct {
		input    string
		re       *regexp.Regexp
		expected Typ
		err      error
	}{
		{
			input: "John 30",
			re:    regexp.MustCompile(`(?P<Name>\w+) (?P<Age>\d+)`),
			expected: Typ{
				Name: "John",
				Age:  30,
			},
		},
		{
			input: "Alice 25",
			re:    regexp.MustCompile(`(?P<Name>\w+) (?P<Age>\d+)`),
			expected: Typ{
				Name: "Alice",
				Age:  25,
			},
		},
	}

	for _, test := range tests {
		result, err := FromString[Typ](test.input, test.re)

		if !errors.Is(err, test.err) {
			t.Errorf("Input: %s\nExpected error: %v\nGot error: %v", test.input, test.err, err)
		}

		if !cmp.Equal(*result, test.expected) {
			t.Errorf("Input: %s\nExpected: %+v\nGot: %+v", test.input, test.expected, result)
		}
	}
}
