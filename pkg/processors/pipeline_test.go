package processors

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPipeline(t *testing.T) {
	type args struct {
		procs []Processor
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should configure a new pipeline instance",
			args: args{
				procs: []Processor{
					&mockProcessor{err: errors.New("some-error")},
					&mockProcessor{err: errors.New("some-other-error")},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPipeline(tt.args.procs...).(*pipeline); !reflect.DeepEqual(got.procs, tt.args.procs) {
				t.Errorf("NewPipeline() = %v, want %v", got.procs, tt.args.procs)
			}
		})
	}
}

func Test_pipeline_ProcessFile(t *testing.T) {
	type fields struct {
		procs []Processor
	}
	type args struct {
		path   string
		reader io.Reader
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		shouldProcess int
		wantErr       error
	}{
		{
			name: "Should only call the first processor if it raises error",
			fields: fields{
				procs: []Processor{
					&mockProcessor{err: errors.New("some-first-error")},
					&mockProcessor{err: errors.New("some-second-error")},
				},
			},
			args: args{
				path:   "some-first-error-path",
				reader: strings.NewReader("some-first-error-reader"),
			},
			shouldProcess: 1,
			wantErr:       errors.New("some-first-error"),
		},
		{
			name: "Should only call the first 2 processors if the second raises error",
			fields: fields{
				procs: []Processor{
					&mockProcessor{},
					&mockProcessor{err: errors.New("some-second-error")},
					&mockProcessor{err: errors.New("some-third-error")},
				},
			},
			args: args{
				path:   "some-second-error-path",
				reader: strings.NewReader("some-second-error-reader"),
			},
			shouldProcess: 2,
			wantErr:       errors.New("some-second-error"),
		},
		{
			name: "Should call all processors if none raises an error",
			fields: fields{
				procs: []Processor{
					&mockProcessor{},
					&mockProcessor{},
					&mockProcessor{},
				},
			},
			args: args{
				path:   "some-no-error-path",
				reader: strings.NewReader("some-no-error-reader"),
			},
			shouldProcess: 3,
			wantErr:       nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pipeline{
				procs: tt.fields.procs,
			}

			err := p.ProcessFile(tt.args.path, tt.args.reader)

			for i := 0; i < tt.shouldProcess; i++ {
				assert.True(t, p.procs[i].(*mockProcessor).processed)
			}
			for i := tt.shouldProcess; i < len(p.procs); i++ {
				assert.False(t, p.procs[i].(*mockProcessor).processed)
			}
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
