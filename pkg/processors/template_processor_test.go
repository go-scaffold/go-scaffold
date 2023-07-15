package processors

import (
	"errors"
	"io"
	"strings"
	"testing"
	"text/template"

	"github.com/pasdam/go-scaffold/pkg/core"
	"github.com/pasdam/go-utils/pkg/assertutils"
	"github.com/stretchr/testify/assert"
)

func TestNewTemplateProcessor(t *testing.T) {
	type args struct {
		data          interface{}
		funcMap       template.FuncMap
		nextProcessor core.Processor
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should create instance with specified fields",
			args: args{
				data:    "some-data",
				funcMap: template.FuncMap{},
				nextProcessor: &templateProcessor{
					data: "some-mock-processor",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTemplateProcessor(tt.args.data, tt.args.nextProcessor, tt.args.funcMap).(*templateProcessor)

			assert.Equal(t, tt.args.data, got.data)
			assert.Equal(t, tt.args.funcMap, got.funcMap)
			assert.Equal(t, tt.args.nextProcessor, got.nextProcessor)
		})
	}
}

func Test_templateProcessor_ProcessFile(t *testing.T) {
	type mocks struct {
		processTemplateErr error
		nextProcessorErr   error
	}
	type args struct {
		filePath string
		reader   io.Reader
	}
	tests := []struct {
		name    string
		mocks   mocks
		args    args
		wantErr error
	}{
		{
			name: "Should propagate error if templates.ProcessTemplate raises it",
			mocks: mocks{
				processTemplateErr: errors.New("some-process-template-error"),
			},
			args: args{
				filePath: "some-process-template-path",
				reader:   strings.NewReader("some-process-template-reader"),
			},
			wantErr: errors.New("some-process-template-error"),
		},
		{
			name: "Should propagate error if next processor raises it",
			mocks: mocks{
				nextProcessorErr: errors.New("some-next-processor-error"),
			},
			args: args{
				filePath: "some-next-processor-path",
				reader:   strings.NewReader("some-next-processor-reader"),
			},
			wantErr: errors.New("some-next-processor-error"),
		},
		{
			name:  "Should not return error if the process succeed",
			mocks: mocks{},
			args: args{
				filePath: "some-success-path",
				reader:   strings.NewReader("some-success-reader"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &templateProcessor{
				data:          tt.name,
				funcMap:       template.FuncMap{},
				nextProcessor: &mockProcessor{err: tt.mocks.nextProcessorErr},
			}
			processedReader := strings.NewReader(tt.name)
			mockTemplatesProcessTemplate(t, tt.args.reader, p.data, p.funcMap, processedReader, tt.mocks.processTemplateErr)

			err := p.ProcessFile(tt.args.filePath, tt.args.reader)

			assertutils.AssertEqualErrors(t, tt.wantErr, err)
		})
	}
}

func mockTemplatesProcessTemplate(t *testing.T, expectedReader io.Reader, expectedData interface{}, expectedFuncMap template.FuncMap, outReader io.Reader, err error) {
	originalValue := templatesProcessTemplate
	templatesProcessTemplate = func(reader io.Reader, data interface{}, funcMap template.FuncMap) (io.Reader, error) {
		assert.Equal(t, expectedReader, reader)
		assert.Equal(t, expectedData, data)
		assert.Equal(t, expectedFuncMap, funcMap)
		return outReader, err
	}
	t.Cleanup(func() { templatesProcessTemplate = originalValue })
}
