// Package imgconv implements image converter
package imgconv

import (
	"io/fs"
	"path/filepath"
)

func isNoEntry(err error) bool {
	return isErrorOccured(err)
}

func isSkip(info fs.DirEntry, path, ext string) bool {
	return info.IsDir() || hasValidFileExtent(path, ext)
}

func validateArg(dirs []string, inExt, outExt string) error {
	if dirs == nil {
		return dirsErr
	}
	if !isValidFileExtent(inExt) || !isValidFileExtent(outExt) {
		return invalidExt
	}
	return nil
}

// Run converts image files that exist in a directory passed as a command line argument.
// The file to be converted is specified by -i.
// The file to be converted is specified by -o as well.
// The image formats supported are jpeg, png, and gif.
// If no image format is specified, jpeg files will be converted to png files.
// Even if the specified directory has subdirectories, image files under the subdirectories will be converted.
// If no directory is passed as an argument, an error will be returned.
// It also returns an error if the appropriate image format is not specified.
// If multiple directories are passed, it will search the directories in the order they are passed.
// Even if a text file or other file not to be converted is found during the search, it will continue to convert other files.
func Run(dirs []string, inExt, outExt string) (convErr error) {
	if err := validateArg(dirs, inExt, outExt); err != nil {
		return err
	}
	for _, entry := range dirs {
		err := filepath.WalkDir(entry, func(path string, info fs.DirEntry, err error) error {
			if isNoEntry(err) {
				return err
			}
			if isSkip(info, path, outExt) {
				return nil
			}
			if !hasValidFileExtent(path, inExt) {
				convErr = wrapErrorWithPath(convErr, path)
				return nil
			}
			if err := convert(path, outExt); err != nil {
				convErr = wrapErrorWithTrim(convErr, err)
				return nil
			}
			return nil
		})
		if err != nil {
			convErr = wrapErrorWithTrim(convErr, err)
		}
	}
	return convErr
}
