package scaffold

import (
	"github.com/pasdam/go-scaffold/pkg/filters"
)

// FileProvider identifies a type used to process files
type FileProvider interface {

	// ProvideFiles provides the files that match filesFilter to the given processor
	ProvideFiles(filesFilter filters.Filter, processor FileProcessor) error
}
