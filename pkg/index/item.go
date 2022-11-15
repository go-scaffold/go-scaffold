package index

import (
	"io/fs"
	"path/filepath"
)

type Item struct {
	info fs.FileInfo
	path string
}

func newItem(parent string, info fs.FileInfo) *Item {
	return &Item{
		info: info,
		path: filepath.Join(parent, info.Name()),
	}
}

func (i *Item) IsDir() bool {
	return i.info.IsDir()
}

func (i *Item) Path() string {
	return i.path
}
