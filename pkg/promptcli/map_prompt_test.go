package promptcli

import (
	"reflect"
	"testing"

	"github.com/manifoldco/promptui"
	"github.com/pasdam/go-scaffold/pkg/prompt"
)

func Test_mapPrompt(t *testing.T) {
	type args struct {
		in *prompt.Entry
	}
	tests := []struct {
		name string
		args args
		want *promptData
	}{
		{
			args: args{
				in: &prompt.Entry{
					Name:    "p1",
					Default: "dp1",
					Type:    "string",
					Message: "message1",
				},
			},
			want: &promptData{
				Prompt: &promptui.Prompt{
					Label:   "message1",
					Default: "dp1",
				},
				Name: "p1",
			},
		},
		{
			args: args{
				in: &prompt.Entry{
					Name:    "p2",
					Default: "33",
					Type:    "int",
					Message: "message2",
				},
			},
			want: &promptData{
				Prompt: &promptui.Prompt{
					Label:    "message2",
					Default:  "33",
					Validate: validateInteger,
				},
				Name: "p2",
			},
		},
		{
			args: args{
				in: &prompt.Entry{
					Name:    "p3",
					Default: "y",
					Type:    "bool",
					Message: "message3",
				},
			},
			want: &promptData{
				Prompt: &promptui.Prompt{
					Label:     "message3",
					Default:   "y",
					IsConfirm: true,
				},
				Name: "p3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapPrompt(tt.args.in)
			gotPrompt := got.Prompt.(*promptui.Prompt)
			wantPrompt := tt.want.Prompt.(*promptui.Prompt)

			if gotPrompt.Label != wantPrompt.Label ||
				gotPrompt.Default != wantPrompt.Default ||
				gotPrompt.IsConfirm != wantPrompt.IsConfirm ||
				got.Name != tt.want.Name ||
				reflect.ValueOf(gotPrompt.Validate).Pointer() != reflect.ValueOf(wantPrompt.Validate).Pointer() {
				t.Errorf("mapPrompt() = %v, want %v", got, tt.want)
			}
		})
	}
}
