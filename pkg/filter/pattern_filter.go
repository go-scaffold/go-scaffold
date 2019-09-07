package filter

import (
	"regexp"
)

type PatternFilter struct {
	patterns  []*regexp.Regexp
	inclusive bool
}

func NewPatternFilter(inclusive bool, patterns ...string) (*PatternFilter, error) {
	regExps := make([]*regexp.Regexp, len(patterns))

	for i := 0; i < len(patterns); i++ {
		regex, err := regexp.Compile(patterns[i])
		if err != nil {
			return nil, err
		}
		regExps[i] = regex
	}

	return &PatternFilter{
		patterns:  regExps,
		inclusive: inclusive,
	}, nil
}

func (f *PatternFilter) Accept(filePath string) bool {
	var valueWhenFound bool
	if f.inclusive {
		valueWhenFound = true

	} else {
		valueWhenFound = false
	}

	for _, pattern := range f.patterns {
		if pattern.MatchString(filePath) {
			// the filter is exclusive but the name matches
			return valueWhenFound
		}
	}
	return !valueWhenFound
}

func (f *PatternFilter) NewInstance(inclusive bool) Filter {
	return &PatternFilter{
		patterns:  f.patterns,
		inclusive: inclusive,
	}
}
