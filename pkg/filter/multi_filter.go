package filter

type multiFilter struct {
	filters []Filter
}

// NewMultiFilter returns a new Filter that merges the input ones,
// it performs a logical OR (if there is one that matches the value
// it will return true)
func NewMultiFilter(filters ...Filter) Filter {
	return &multiFilter{
		filters: filters,
	}
}

func (f *multiFilter) Accept(value string) bool {
	for _, filter := range f.filters {
		if filter.Accept(value) {
			return true
		}
	}
	return false
}
