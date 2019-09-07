package filter

type Filter interface {
	Accept(filePath string) bool
}
