package strunpack

import "regexp"

func FromString[T interface{}](s string, re *regexp.Regexp) (*T, error) {
	var result T
	err := Unpack(s, re, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
