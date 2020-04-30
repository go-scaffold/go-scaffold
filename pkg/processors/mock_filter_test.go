package processors

type mockFilter struct {
	accept bool
}

func (f *mockFilter) Accept(_ string) bool {
	return f.accept
}
