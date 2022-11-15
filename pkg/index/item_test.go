package index

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newItem(t *testing.T) {
	testdataChildren, err := ioutil.ReadDir("testdata")
	assert.NoError(t, err)
	assert.Len(t, testdataChildren, 4)

	type args struct {
		parent string
		info   fs.FileInfo
	}
	type wantItem struct {
		isDir bool
		path  string
	}
	tests := []struct {
		name string
		args args
		want *wantItem
	}{
		{
			name: "Should correctly wrap a directory",
			args: args{
				parent: "sub1",
				info:   testdataChildren[2],
			},
			want: &wantItem{
				isDir: true,
				path:  filepath.Join("sub1", testdataChildren[2].Name()),
			},
		},
		{
			name: "Should correctly wrap a file",
			args: args{
				parent: "sub2",
				info:   testdataChildren[0],
			},
			want: &wantItem{
				isDir: false,
				path:  filepath.Join("sub2", testdataChildren[0].Name()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newItem(tt.args.parent, tt.args.info)

			assert.Equal(t, tt.want.isDir, got.IsDir())
			assert.Equal(t, tt.want.path, got.Path())
		})
	}
}
