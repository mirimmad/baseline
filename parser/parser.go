package parser

import (
	"errors"
	"regexp"
)

type Source struct {
	str   string
	index int
}

type ParserResult struct {
	value  interface{}
	source Source
}

func (s *Source) match(pattern string) (*ParserResult, error) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return nil, errors.New("Regex Error")
	}
	match_ := r.FindStringIndex(s.str[s.index:])
	if len(match_) == 0 {
		return nil, errors.New("No Match")
	}
	newIndex := s.index + (match_[1] - match_[0])
	source := Source{s.str, newIndex}
	return &ParserResult{match_, source}, nil
}
