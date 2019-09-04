package scaffold

import (
	"regexp"
)

type patternFilter struct {
	patterns []*regexp.Regexp
}

func NewPatternFilter(patterns ...string) (Filter, error) {
	regExps := make([]*regexp.Regexp, len(patterns))

	for i := 0; i < len(patterns); i++ {
		regex, err := regexp.Compile(patterns[i])
		if err != nil {
			return nil, err
		}
		regExps[i] = regex
	}

	return &patternFilter{
		patterns: regExps,
	}, nil
}

func (f *patternFilter) Accept(filePath string) bool {
	for _, pattern := range f.patterns {
		if pattern.MatchString(filePath) {
			return false
		}
	}
	return true
}
