package index

import (
	"io"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/pasdam/go-utils/pkg/assertutils"
	"github.com/stretchr/testify/assert"
)

func Test_index_Prepend(t *testing.T) {
	testdataChildren, err := ioutil.ReadDir("testdata")
	assert.NoError(t, err)
	subfolderChildren, err := ioutil.ReadDir(filepath.Join("testdata", "sub1"))
	assert.NoError(t, err)

	testdataItems := itemsList("", testdataChildren)
	subfolderItems := itemsList("", subfolderChildren)

	type fields struct {
		items []*Item
	}
	type args struct {
		items []*Item
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   fields
	}{
		{
			name: "Should have subfolder's children first",
			fields: fields{
				items: testdataItems,
			},
			args: args{
				items: subfolderItems,
			},
			want: fields{
				items: append(subfolderItems, testdataItems...),
			},
		},
		{
			name: "Should have subfolder's children last",
			fields: fields{
				items: subfolderItems,
			},
			args: args{
				items: testdataItems,
			},
			want: fields{
				items: append(testdataItems, subfolderItems...),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &index{
				items: tt.fields.items,
			}

			i.Prepend(tt.args.items...)

			if !reflect.DeepEqual(i.items, tt.want.items) {
				t.Errorf("index.Prepend() = %v, want %v", i.items, tt.want.items)
			}
		})
	}
}

func Test_index_TakeFirst(t *testing.T) {
	testdataChildren, err := ioutil.ReadDir("testdata")
	assert.NoError(t, err)
	assert.Len(t, testdataChildren, 4)

	testdataItems := itemsList("", testdataChildren)

	type fields struct {
		items []*Item
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Item
		wantErr error
	}{
		{
			name: "Should return EOF if items list is nil",
			fields: fields{
				items: nil,
			},
			want:    nil,
			wantErr: io.EOF,
		},
		{
			name: "Should return EOF if items list is empty",
			fields: fields{
				items: []*Item{},
			},
			want:    nil,
			wantErr: io.EOF,
		},
		{
			name: "Should return first element in the list",
			fields: fields{
				items: testdataItems,
			},
			want:    testdataItems[0],
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &index{
				items: tt.fields.items,
			}

			got, err := i.TakeFirst()

			assertutils.AssertEqualErrors(t, tt.wantErr, err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("index.TakeFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
