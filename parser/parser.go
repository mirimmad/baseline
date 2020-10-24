package parser

import (
	"errors"
	"regexp"
)

type ParseFn func(s Source) *ParserResult

type Source struct {
	Str   string
	Index int
}

type ParserResult struct {
	value  interface{}
	source Source
}

func (s *Source) Match(pattern string) (*ParserResult, error) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return nil, errors.New("Regex Error")
	}
	match_ := r.FindStringIndex(s.Str[s.Index:])
	if len(match_) == 0 {
		return nil, errors.New("No Match")
	}
	newIndex := s.Index + (match_[1] - match_[0])
	source := Source{s.Str, newIndex}
	return &ParserResult{match_, source}, nil
}

type Parser struct {
	Parse ParseFn
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Regex(pattern string) *Parser {
	return &Parser{func(s Source) *ParserResult {
		r, err := s.Match(pattern)
		if err != nil {
			return nil
		}
		return r
	}}
}

func (p *Parser) Constant(value interface{}) *Parser {
	return &Parser{func(s Source) *ParserResult {
		return &ParserResult{value, s}
	}}
}

func (p1 *Parser) Or(p2 *Parser) *Parser {
	return &Parser{func(s Source) *ParserResult {
		result := p1.Parse(s)
		if result != nil {
			return result
		}
		return p2.Parse(s)
	}}
}

// func (t *Parser) ZeroOrMore(p *Parser) *Parser {
// 	return &Parser{func(s Source) *ParserResult {
// 		results := make([]ParserResult, 0)

// 	}}
// }
