package index

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_itemsList(t *testing.T) {
	testdataChildren, err := ioutil.ReadDir("testdata")
	assert.NoError(t, err)
	subfolderChildren, err := ioutil.ReadDir(filepath.Join("testdata", "sub1"))
	assert.NoError(t, err)

	type args struct {
		parent string
		items  []fs.FileInfo
	}
	type wantItem struct {
		isDir bool
		path  string
	}

	tests := []struct {
		name string
		args args
		want []*wantItem
	}{
		{
			name: "Should create a list for testdata folder",
			args: args{
				parent: "",
				items:  testdataChildren,
			},
			want: []*wantItem{
				{false, "a.txt"},
				{false, "b.txt"},
				{true, "sub1"},
				{false, "t.txt"},
			},
		},
		{
			name: "Should create a list for testdata subfolder",
			args: args{
				parent: "testdata",
				items:  subfolderChildren,
			},
			want: []*wantItem{
				{false, filepath.Join("testdata", "a1.txt")},
				{false, filepath.Join("testdata", "b1.txt")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := itemsList(tt.args.parent, tt.args.items)

			assert.Len(t, got, len(tt.want))
			for i := 0; i < len(tt.want); i++ {
				assert.Equal(t, tt.want[i].isDir, got[i].IsDir(), "%d", i)
				assert.Equal(t, tt.want[i].path, got[i].Path(), "%d", i)
			}
		})
	}
}
