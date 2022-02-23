package imgconv_test

import (
	"errors"
	"os"
	"testing"

	"example.com/ex01/imgconv"
)

func TestConvertImage(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		want error
	}{
		{"normal", []string{"../convert", "../testdata"}, nil},
		{"jpg to png", []string{"../convert", "-i=jpg", "-o=png", "../testdata"}, nil},
		{"jpg to gif", []string{"../convert", "-i=jpg", "-o=gif", "../testdata"}, nil},
		{"png to jpg", []string{"../convert", "-i=png", "-o=jpg", "../testdata"}, nil},
		{"png to gif", []string{"../convert", "-i=png", "-o=gif", "../testdata"}, nil},
		{"gif to jpg", []string{"../convert", "-i=gif", "-o=jpg", "../testdata"}, nil},
		{"gif to png", []string{"../convert", "-i=gif", "-o=png", "../testdata"}, nil},
		{"jpg to jpg", []string{"../convert", "-i=jpg", "-o=jpg", "../testdata"}, nil},
		{"png to png", []string{"../convert", "-i=png", "-o=png", "../testdata"}, nil},
		{"gif to gif", []string{"../convert", "-i=gif", "-o=gif", "../testdata"}, nil},
		{"no such dir", []string{"../convert", "./hoge"}, nil},
		{"no such dir", []string{"../convert", "-i=jpg", "-o=png", "./hoge"}, nil},
		{"no such dir", []string{"../convert", "-i=jpg", "-o=gif", "./hoge"}, nil},
		{"no such dir", []string{"../convert", "-i=png", "-o=jpg", "./hoge"}, nil},
		{"no such dir", []string{"../convert", "-i=png", "-o=gif", "./hoge"}, nil},
		{"no such dir", []string{"../convert", "-i=gif", "-o=jpg", "./hoge"}, nil},
		{"no such dir", []string{"../convert", "-i=gif", "-o=png", "./hoge"}, nil},
		{"no argument", []string{"../convert"}, errors.New("error: invalid argument")},
		{"invalid option", []string{"../convert", "-i=hoge", "-o=huga", "../testdata"}, errors.New("error: invalid extension")},
		{"invalid option", []string{"../convert", "-i=png", "-o=txt", "../testdata"}, errors.New("error: invalid extension")},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			os.Args = tt.arg
			res := imgconv.ConvertImage()
			assertError(t, tt.want, res)
		})
	}
}

func assertError(t *testing.T, expect, got error) {
	t.Helper()
	if expect == nil && got == nil {
		return
	}
	if expect != nil && got == nil {
		t.Errorf("expected: %v, got: nil", expect)
	} else if expect == nil && got != nil {
		t.Errorf("expected: nil, got: %v", got)
	} else if expect.Error() != got.Error() {
		t.Errorf("expected: %v, got: %v", expect, got)
	}
}
