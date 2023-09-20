package strunpack

import "regexp"

type Unpacker[T any] struct {
	re *regexp.Regexp
}

func FromRegex[T any](re *regexp.Regexp) *Unpacker[T] {
	return &Unpacker[T]{re}
}

func FromString[T any](re string) *Unpacker[T] {
	return &Unpacker[T]{regexp.MustCompile(re)}
}

func FromStringWithError[T any](reStr string) (*Unpacker[T], error) {
	re, err := regexp.Compile(reStr)
	if err != nil {
		return nil, err
	}
	return &Unpacker[T]{re}, nil
}

func (u *Unpacker[T]) Unpack(s string) (*T, error) {
	res, err := fromString[T](s, u.re)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func fromString[T interface{}](s string, re *regexp.Regexp) (*T, error) {
	var result T
	err := Unpack(s, re, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
