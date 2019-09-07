package filter

import (
	"regexp"
)

type patternFilter struct {
	patterns  []*regexp.Regexp
	inclusive bool
}

func NewPatternFilter(inclusive bool, patterns ...string) (Filter, error) {
	regExps := make([]*regexp.Regexp, len(patterns))

	for i := 0; i < len(patterns); i++ {
		regex, err := regexp.Compile(patterns[i])
		if err != nil {
			return nil, err
		}
		regExps[i] = regex
	}

	return &patternFilter{
		patterns:  regExps,
		inclusive: inclusive,
	}, nil
}

func (f *patternFilter) Accept(filePath string) bool {
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
