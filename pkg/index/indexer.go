package index

import (
	"io"
	"io/ioutil"
	"path/filepath"
)

type Indexer struct {
	Dir string

	children *index
}

func (i *Indexer) index(abs, relative string) error {
	filesInfo, err := ioutil.ReadDir(abs)
	if err != nil {
		return err
	}
	i.children.Prepend(itemsList(relative, filesInfo)...)
	return nil
}

func (i *Indexer) NextFile() (*Item, error) {
	if i.children == nil {
		i.children = &index{}
		err := i.index(i.Dir, "")
		if err != nil {
			return nil, err
		}
	}

	firstChild, err := i.children.TakeFirst()
	for ; err == nil && firstChild.IsDir(); firstChild, err = i.children.TakeFirst() {
		i.index(filepath.Join(i.Dir, firstChild.Path()), firstChild.Path())
	}

	if firstChild == nil {
		return nil, io.EOF
	}

	return firstChild, nil
}
