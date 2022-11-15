package index

import (
	"errors"
	"io"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/pasdam/mockit/mockit"
)

func TestIndexer_NextFile(t *testing.T) {
	type mocks struct {
		readDir error
	}
	type fields struct {
		Dir      string
		children *index
	}
	tests := []struct {
		name    string
		fields  fields
		mocks   mocks
		want    []string
		wantErr error
	}{
		{
			name: "Should index dir and return elements in the right order",
			fields: fields{
				children: nil,
				Dir:      "./testdata",
			},
			mocks: mocks{
				readDir: nil,
			},
			want: []string{
				"a.txt",
				"b.txt",
				filepath.Join("sub1", "a1.txt"),
				filepath.Join("sub1", "b1.txt"),
				"t.txt",
			},
			wantErr: io.EOF,
		},
		{
			name: "Should propagate error if ReadDir raises it",
			fields: fields{
				children: nil,
				Dir:      "./testdata",
			},
			mocks: mocks{
				readDir: errors.New("some ReadDir error"),
			},
			want:    nil,
			wantErr: errors.New("some ReadDir error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indexer := &Indexer{
				Dir:      tt.fields.Dir,
				children: tt.fields.children,
			}

			if tt.mocks.readDir != nil {
				mockFunc := mockit.MockFunc(t, ioutil.ReadDir)
				mockFunc.With(tt.fields.Dir).Return(nil, tt.mocks.readDir)
			}

			got, err := indexer.NextFile()

			for i := 0; i < len(tt.want); i++ {
				if !reflect.DeepEqual(got.Path(), tt.want[i]) {
					t.Errorf("Indexer.NextFile() = %v, want %v", got.Path(), tt.want[i])
				}

				got, err = indexer.NextFile()
			}

			assertEqualErrors(t, tt.wantErr, err)
		})
	}
}
