package index

import (
	"io"
)

type index struct {
	items []*Item
}

func (i *index) Prepend(items ...*Item) {
	i.items = append(items, i.items...)
}

func (i *index) TakeFirst() (*Item, error) {
	if len(i.items) == 0 {
		return nil, io.EOF
	}
	result := i.items[0]
	i.items = i.items[1:]
	return result, nil
}
