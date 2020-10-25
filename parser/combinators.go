package parser

import (
	"errors"
	"regexp"
)

type ParseFn func(s Source) *ParserResult
type CallBack func(value interface{}) *Parser
type CallBackForMap func(value interface{}) interface{}

type Source struct {
	Str   string
	Index int
}

type ParserResult struct {
	Value   interface{}
	Source_ Source
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

func (t *Parser) ZeroOrMore(p *Parser) *Parser {
	return &Parser{func(s Source) *ParserResult {
		results := make([]ParserResult, 0)
		item := p.Parse(s)
		for {

			if item == nil {
				break
			}
			results = append(results, *item)
			s = item.Source_
			item = p.Parse(s)
		}
		return &ParserResult{results, s}
	}}
}

func (p *Parser) Bind(callback CallBack) *Parser {
	return &Parser{func(s Source) *ParserResult {
		result := p.Parse(s)
		if result != nil {
			value := result.Value
			source := result.Source_
			return callback(value).Parse(source)
		}
		return nil
	}}
}

func (p *Parser) And(t *Parser) *Parser {
	return p.Bind(func(_ interface{}) *Parser {
		return t
	})
}

func (p *Parser) Map(callback CallBackForMap) *Parser {
	return p.Bind(func(value interface{}) *Parser {
		return p.Constant(callback(value))
	})
}

func (t *Parser) Maybe(p *Parser) *Parser {
	return t.Or(p.Constant(nil))
}

func (p *Parser) ParseStringToCompletion(str string) interface{} {
	source := Source{str, 0}
	result := p.Parse(source)
	if result == nil {
		panic("Parse error at index 0.")
	}
	index := result.Source_.Index
	if index != len(result.Source_.Str) {
		panic("Parse error at index __")
	}
	return result.Value
}
