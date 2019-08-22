package promptcli

import (
	"reflect"
	"testing"

	"github.com/manifoldco/promptui"
)

func Test_validateFuncForType(t *testing.T) {
	type args struct {
		promptType string
	}
	tests := []struct {
		name string
		args args
		want promptui.ValidateFunc
	}{
		{ args: args{ promptType:"int" }, want: validateInteger},
		{ args: args{ promptType:"string" }, want: nil},
		{ args: args{ promptType:"bool" }, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateFuncForType(tt.args.promptType); reflect.ValueOf(got).Pointer() != reflect.ValueOf(tt.want).Pointer() {
				t.Errorf("validateFuncForType() = %v, want %v", got, tt.want)
			}
		})
	}
}
