package scaffold

import (
	"github.com/pasdam/go-scaffold/pkg/filter"
)

type FileProvider interface {
	ProvideFiles(filesFilter filter.Filter, processor FileProcessor) error
}
