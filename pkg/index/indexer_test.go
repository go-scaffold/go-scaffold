package index

import (
	"errors"
	"io"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/pasdam/go-utils/pkg/assertutils"
)

func TestIndexer_NextFile(t *testing.T) {
	type fields struct {
		Dir      string
		children *index
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr error
	}{
		{
			name: "Should index dir and return elements in the right order",
			fields: fields{
				children: nil,
				Dir:      "./testdata",
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
				Dir:      "",
			},
			want:    nil,
			wantErr: errors.New("open : no such file or directory"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indexer := &Indexer{
				Dir:      tt.fields.Dir,
				children: tt.fields.children,
			}

			got, err := indexer.NextFile()

			for i := 0; i < len(tt.want); i++ {
				if !reflect.DeepEqual(got.Path(), tt.want[i]) {
					t.Errorf("Indexer.NextFile() = %v, want %v", got.Path(), tt.want[i])
				}

				got, err = indexer.NextFile()
			}

			assertutils.AssertEqualErrors(t, tt.wantErr, err)
		})
	}
}
