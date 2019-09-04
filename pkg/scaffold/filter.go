package scaffold

type Filter interface {
	Accept(filePath string) bool
}
