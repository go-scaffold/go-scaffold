package scaffold

import (
	"github.com/pasdam/go-scaffold/pkg/filter"
)

// FileProvider identifies a type used to process files
type FileProvider interface {

	// ProvideFiles provides the files that match filesFilter to the given processor
	ProvideFiles(filesFilter filter.Filter, processor FileProcessor) error
}
