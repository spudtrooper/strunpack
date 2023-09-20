package strunpack

import "regexp"

type Unpacker[T any] struct {
	re *regexp.Regexp
}

func UnpackerFromRegex[T any](re *regexp.Regexp) *Unpacker[T] {
	return &Unpacker[T]{re}
}

func UnpackerFromString[T any](re string) *Unpacker[T] {
	return &Unpacker[T]{regexp.MustCompile(re)}
}

func UnpackerFromStringWithError[T any](reStr string) (*Unpacker[T], error) {
	re, err := regexp.Compile(reStr)
	if err != nil {
		return nil, err
	}
	return &Unpacker[T]{re}, nil
}

func (u *Unpacker[T]) Unpack(s string) (*T, error) {
	res, err := FromString[T](s, u.re)
	if err != nil {
		return nil, err
	}
	return res, nil
}
