package app

import (
	"errors"
	"os"
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
			name: "Should return nil if the script does not exist",
			args: args{
				path:    "some-path",
				workDir: "some-work-dir",
			},
			wantErr: nil,
		},
		{
			name: "Should return error raised by cmd.Run",
			args: args{
				path:    "run_init_script_test.go",
				workDir: "./",
			},
			wantErr: errors.New("some-error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called := false
			errHandler := func(v ...interface{}) {
				called = true
				assert.Equal(t, 2, len(v))
				assert.Equal(t, "Error while executing init script. ", v[0])
				assert.Equal(t, tt.wantErr, v[1])
			}
			guard := monkey.PatchInstanceMethod(reflect.TypeOf(&exec.Cmd{}), "Run", func(c *exec.Cmd) error {
				assert.Equal(t, tt.args.workDir, c.Dir)
				assert.Equal(t, tt.args.path, c.Path)
				assert.Equal(t, os.Stderr, c.Stderr)
				assert.Equal(t, os.Stdout, c.Stdout)
				return tt.wantErr
			})
			defer guard.Unpatch()

			runInitScript(tt.args.path, tt.args.workDir, errHandler)

			assert.Equal(t, tt.wantErr != nil, called)
		})
	}
}
