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
			if tt.want == nil && res != nil {
				t.Errorf("imgconv.ConvertImage(): expected: nil, got: %v", res)
			} else if tt.want != nil && res == nil {
				t.Errorf("imgconv.ConvertImage(): expected: %v, got: nil", tt.want)
			} else if (tt.want != nil && res != nil) && (tt.want.Error() != res.Error()) {
				t.Errorf("imgconv.ConvertImage(): expected: %v, got: %v", tt.want, res)
			}
		})
	}
}
