package imgconv_test

import (
	"flag"
	"os"
	"reflect"
	"testing"

	"example.com/ex01/imgconv"
)

func TestValidateArgs(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		iFlag    string
		oFlag    string
		wantDirs []string
		wantFrom string
		wantTo   string
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"default", []string{"cmd", "/some/directory"}, "", "", []string{"/some/directory"}, "jpg", "png", false},
		{"assign format", []string{"cmd", "/some/directory"}, "png", "jpg", []string{"/some/directory"}, "png", "jpg", false},
		{"assign format", []string{"cmd", "/some/directory"}, "png", "jpeg", []string{"/some/directory"}, "png", "jpeg", false},
		{"assign format", []string{"cmd", "/some/directory"}, "png", "gif", []string{"/some/directory"}, "png", "gif", false},
		{"assign format", []string{"cmd", "/some/directory"}, "png", "png", []string{"/some/directory"}, "png", "png", false},
		{"assign format", []string{"cmd", "/some/directory"}, "jpg", "png", []string{"/some/directory"}, "jpg", "png", false},
		{"assign format", []string{"cmd", "/some/directory"}, "jpg", "jpeg", []string{"/some/directory"}, "jpg", "jpeg", false},
		{"assign format", []string{"cmd", "/some/directory"}, "jpg", "gif", []string{"/some/directory"}, "jpg", "gif", false},
		{"assign format", []string{"cmd", "/some/directory"}, "jpg", "jpg", []string{"/some/directory"}, "jpg", "jpg", false},
		{"assign format", []string{"cmd", "/some/directory"}, "jpeg", "gif", []string{"/some/directory"}, "jpeg", "gif", false},
		{"assign format", []string{"cmd", "/some/directory"}, "jpeg", "png", []string{"/some/directory"}, "jpeg", "png", false},
		{"assign format", []string{"cmd", "/some/directory"}, "jpeg", "jpg", []string{"/some/directory"}, "jpeg", "jpg", false},
		{"assign format", []string{"cmd", "/some/directory"}, "jpeg", "jpeg", []string{"/some/directory"}, "jpeg", "jpeg", false},
		{"assign format", []string{"cmd", "/some/directory"}, "gif", "jpg", []string{"/some/directory"}, "gif", "jpg", false},
		{"assign format", []string{"cmd", "/some/directory"}, "gif", "png", []string{"/some/directory"}, "gif", "png", false},
		{"assign format", []string{"cmd", "/some/directory"}, "gif", "jpeg", []string{"/some/directory"}, "gif", "jpeg", false},
		{"assign format", []string{"cmd", "/some/directory"}, "gif", "gif", []string{"/some/directory"}, "gif", "gif", false},
		{"no args", []string{"cmd"}, "", "", nil, "", "", true},
		{"invalid format", []string{"cmd", "/some/directory"}, "txt", "bmp", nil, "", "", true},
		{"invalid format", []string{"cmd", "/some/directory"}, "jpg", "txt", nil, "", "", true},
		{"invalid format", []string{"cmd", "/some/directory"}, "mp3", "mp4", nil, "", "", true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			os.Args = tt.args
			if tt.iFlag != "" {
				flag.CommandLine.Set("i", tt.iFlag)
			}
			if tt.oFlag != "" {
				flag.CommandLine.Set("o", tt.oFlag)
			}
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
