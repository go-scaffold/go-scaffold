package filter

import (
	"regexp"
)

// PatternFilter if a Filter that accept a string if it matches one or more regEx (at least one)
type PatternFilter struct {
	patterns  []*regexp.Regexp
	inclusive bool
}

// NewPatternFilter returns a new instance of PatternFilter configured with the specified regExp
// or an error if one of them is invalid.
// The filter accept a string if it is inclusive and the value matches one of the regexp, of if
// it is exclusive (inclusive=false) and the value doesn't match any of the patterns.
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

// Accept returns true if it is inclusive and the value matches one of the regexp, of if
// it is exclusive (inclusive=false) and the value doesn't match any of the patterns.
func (f *PatternFilter) Accept(value string) bool {
	var valueWhenFound bool
	if f.inclusive {
		valueWhenFound = true

	} else {
		valueWhenFound = false
	}

	for _, pattern := range f.patterns {
		if pattern.MatchString(value) {
			// the filter is exclusive but the name matches
			return valueWhenFound
		}
	}
	return !valueWhenFound
}

// NewInstance creates a new instance of the filter
func (f *PatternFilter) NewInstance(inclusive bool) Filter {
	return &PatternFilter{
		patterns:  f.patterns,
		inclusive: inclusive,
	}
}
