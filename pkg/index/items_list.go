package index

import "io/fs"

func itemsList(parent string, items []fs.FileInfo) []*Item {
	result := make([]*Item, 0, len(items))
	for _, info := range items {
		result = append(result, newItem(parent, info))
	}
	return result
}
