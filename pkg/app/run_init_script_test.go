package app

import (
	"errors"
	"os/exec"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func Test_runInitScript(t *testing.T) {
	type args struct {
		path    string
		workDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Should return error raised by cmd.Run",
			args: args{
				path:    "some-path",
				workDir: "some-work-dir",
			},
			wantErr: errors.New("some-error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			guard := monkey.PatchInstanceMethod(reflect.TypeOf(&exec.Cmd{}), "Run", func(c *exec.Cmd) error {
				assert.Equal(t, tt.args.workDir, c.Dir)
				assert.Equal(t, tt.args.path, c.Path)
				return tt.wantErr
			})
			defer guard.Unpatch()

			if err := runInitScript(tt.args.path, tt.args.workDir); err != tt.wantErr {
				t.Errorf("runInitScript() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
