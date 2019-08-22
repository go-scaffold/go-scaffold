package promptcli

import "testing"

func Test_validateInteger(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{ args: args { value: "invalid-number" }, wantErr: true},
		{ args: args { value: "1" }, wantErr: false},
		{ args: args { value: "0" }, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateInteger(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("validateInteger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
