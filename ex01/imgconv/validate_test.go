package imgconv_test

import (
	"os"
	"reflect"
	"testing"

	"example.com/ex01/imgconv"
)

func TestValidateArgs(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantDirs []string
		wantFrom string
		wantTo   string
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"no args", []string{"cmd"}, nil, "", "", true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			os.Args = tt.args
			gotDirs, gotFrom, gotTo, err := imgconv.ValidateArgs()
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDirs, tt.wantDirs) {
				t.Errorf("ValidateArgs() gotDirs = %v, want %v", gotDirs, tt.wantDirs)
			}
			if gotFrom != tt.wantFrom {
				t.Errorf("ValidateArgs() gotFrom = %v, want %v", gotFrom, tt.wantFrom)
			}
			if gotTo != tt.wantTo {
				t.Errorf("ValidateArgs() gotTo = %v, want %v", gotTo, tt.wantTo)
			}
		})
	}
}
