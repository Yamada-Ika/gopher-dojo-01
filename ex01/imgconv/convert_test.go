package imgconv_test

import (
	"errors"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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
		// {"jpg to png", []string{"../convert", "-i=jpg", "-o=png", "../testdata"}, nil},
		// {"jpg to gif", []string{"../convert", "-i=jpg", "-o=gif", "../testdata"}, nil},
		// {"png to jpg", []string{"../convert", "-i=png", "-o=jpg", "../testdata"}, nil},
		// {"png to gif", []string{"../convert", "-i=png", "-o=gif", "../testdata"}, nil},
		// {"gif to jpg", []string{"../convert", "-i=gif", "-o=jpg", "../testdata"}, nil},
		// {"gif to png", []string{"../convert", "-i=gif", "-o=png", "../testdata"}, nil},
		// {"jpg to jpg", []string{"../convert", "-i=jpg", "-o=jpg", "../testdata"}, nil},
		// {"png to png", []string{"../convert", "-i=png", "-o=png", "../testdata"}, nil},
		// {"gif to gif", []string{"../convert", "-i=gif", "-o=gif", "../testdata"}, nil},
		// {"no such dir", []string{"../convert", "./hoge"}, nil},
		// {"no such dir", []string{"../convert", "-i=jpg", "-o=png", "./hoge"}, nil},
		// {"no such dir", []string{"../convert", "-i=jpg", "-o=gif", "./hoge"}, nil},
		// {"no such dir", []string{"../convert", "-i=png", "-o=jpg", "./hoge"}, nil},
		// {"no such dir", []string{"../convert", "-i=png", "-o=gif", "./hoge"}, nil},
		// {"no such dir", []string{"../convert", "-i=gif", "-o=jpg", "./hoge"}, nil},
		// {"no such dir", []string{"../convert", "-i=gif", "-o=png", "./hoge"}, nil},
		// {"no argument", []string{"../convert"}, errors.New("error: invalid argument")},
		// {"invalid option", []string{"../convert", "-i=hoge", "-o=huga", "../testdata"}, errors.New("error: invalid extension")},
		// {"invalid option", []string{"../convert", "-i=png", "-o=txt", "../testdata"}, errors.New("error: invalid extension")},
	}
	// With image format test
	// tt := tests[0]
	// t.Run(tt.name, func(t *testing.T) {
	// 	os.Args = tt.arg
	// 	res := imgconv.ConvertImage()
	// 	assertError(t, tt.want, res)
	// 	assertDirStruct(t, tt.arg[len(tt.arg)-1], 0)
	// })

	// Without image format teset
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

var targetImages = [][]string{
	{
		"../testdata/sub_dir1/.jpeg.png",
		"../testdata/sub_dir1/test2.png",
		"../testdata/sub_dir1/.jpg.png",
		"../testdata/sub_dir1/test4.png",
		"../testdata/.jpeg.png",
		"../testdata/test2.png",
		"../testdata/.jpg.png",
		"../testdata/test4.png",
	},
}

var expectedDirStruct = []string{
	"../testdata",
	"../testdata/.gif.gif.gif",
	"../testdata/.DS_Store",
	"../testdata/test6.jpg",
	"../testdata/sub_dir1",
	"../testdata/sub_dir1/.gif.gif.gif",
	"../testdata/sub_dir1/.DS_Store",
	"../testdata/sub_dir1/test6.jpg",
	"../testdata/sub_dir1/.jpg.jpg.jpg",
	"../testdata/sub_dir1/test7.png",
	"../testdata/sub_dir1/.jpeg.jpeg",
	"../testdata/sub_dir1/test3.png",
	"../testdata/sub_dir1/test2.jpg",
	"../testdata/sub_dir1/.jpg.jpg",
	"../testdata/sub_dir1/.png.png",
	"../testdata/sub_dir1/.gif.gif",
	"../testdata/sub_dir1/.jpeg.jpeg.jpeg",
	"../testdata/sub_dir1/.png.png.png",
	"../testdata/sub_dir1/test8.jpeg",
	"../testdata/sub_dir1/test5.gif",
	"../testdata/sub_dir1/test4.jpeg",
	"../testdata/sub_dir1/test1.gif",
	"../testdata/.jpg.jpg.jpg",
	"../testdata/test7.png",
	"../testdata/.jpeg.jpeg",
	"../testdata/test3.png",
	"../testdata/test2.jpg",
	"../testdata/.jpg.jpg",
	"../testdata/.png.png",
	"../testdata/.gif.gif",
	"../testdata/.png.png.png",
	"../testdata/test5.gif",
	"../testdata/test4.jpeg",
	"../testdata/test1.gif",
}

func isValidFileExtent(path string) bool {
	return strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg") || strings.HasSuffix(path, ".png") || strings.HasSuffix(path, ".gif")
}

func validateImageFormat(image string) error {
	f, err := os.Open(image)
	if err != nil {
		return err
	}
	if strings.HasSuffix(image, ".jpg") || strings.HasSuffix(image, ".jpeg") {
		_, err = jpeg.Decode(f)
		if err != nil {
			return errors.New("unmatch file extension and file type")
		}
	} else if strings.HasSuffix(image, ".png") {
		_, err = png.Decode(f)
		if err != nil {
			return errors.New("unmatch file extension and file type")
		}
	} else if strings.HasSuffix(image, ".gif") {
		_, err = gif.Decode(f)
		if err != nil {
			return errors.New("unmatch file extension and file type")
		}
	}
	defer f.Close()
	defer os.Remove(image)
	return nil
}

func validateAllImageFormat(targetImages []string, resultDirStruct map[string]int) error {
	for _, image := range targetImages {
		if _, ok := resultDirStruct[image]; ok {
			if isValidFileExtent(image) {
				if err := validateImageFormat(image); err != nil {
					return err
				}
			}
		} else {
			return errors.New("image file does not exist")
		}
	}
	return nil
}

func assertDirStruct(t *testing.T, testDir string, index int) {
	t.Helper()
	res := make(map[string]int)
	filepath.WalkDir(testDir, func(path string, info fs.DirEntry, err error) error {
		if _, ok := res[path]; !ok {
			res[path] = 1
		}
		return nil
	})

	if err := validateAllImageFormat(targetImages[index], res); err != nil {
		t.Errorf("error: %v", err)
		return
	}

	for _, image := range expectedDirStruct {
		if _, ok := res[image]; !ok {
			t.Errorf("error")
			return
		}
	}
}
