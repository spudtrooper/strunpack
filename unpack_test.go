package strunpack

import (
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
)

type Typ struct {
	Name string
	Age  int
}

func TestUnpackValid(t *testing.T) {
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
		var result Typ
		err := Unpack(test.input, test.re, &result)

		if !cmp.Equal(result, test.expected) {
			t.Errorf("Input: %s\nExpected: %+v\nGot: %+v", test.input, test.expected, result)
		}

		if !errors.Is(err, test.err) {
			t.Errorf("Input: %s\nExpected error: %v\nGot error: %v", test.input, test.err, err)
		}
	}
}

func TestUnpackInvalid(t *testing.T) {
	tests := []struct {
		input string
		re    *regexp.Regexp
	}{
		{
			input: "InvalidInput",
			re:    regexp.MustCompile(`(?P<Name>\w+) (?P<Age>\d+)`),
		},
		{
			input: "NilRE",
		},
		{
			input: "Bob twenty",
			re:    regexp.MustCompile(`(?P<Name>\w+) (?P<Age>\w+)`),
		},
	}

	for _, test := range tests {
		var result Typ
		err := Unpack(test.input, test.re, &result)

		if err == nil {
			t.Errorf("Expecting error:\nInput: %s\nGot error: %v", test.input, err)
		}
	}
}

func TestUnpackInvalidResultType(t *testing.T) {
	input := "Jane 20"
	re := regexp.MustCompile(`(?P<Name>\w+) (?P<Age>\d+)`)
	var result Typ
	err := Unpack(input, re, result)

	if err == nil {
		t.Errorf("Expecting error:\nInput: %s\nExpected error: %v\nGot error: %v", input, err, err)
	}
	if want, have := "Invalid result type. Expected a pointer to a struct", err.Error(); want != have {
		t.Errorf("Mismatched strings:\nInput: %s\nExpected error: %v\nGot error: %v", input, want, have)
	}
}
