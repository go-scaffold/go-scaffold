package filter

type multiFilter struct {
	filters []Filter
}

func NewMultiFilter(filters ...Filter) Filter {
	return &multiFilter{
		filters: filters,
	}
}

func (f *multiFilter) Accept(filePath string) bool {
	for _, filter := range f.filters {
		if filter.Accept(filePath) {
			return true
		}
	}
	return false
}
