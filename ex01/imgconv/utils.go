package imgconv

import (
	"path/filepath"
	"strings"
)

func isValidFileExtent(ext string) bool {
	switch ext {
	case "jpg", "jpeg", "png", "gif":
		break
	default:
		return false
	}
	return true
}

func hasValidFileExtent(path string, ext string) bool {
	switch ext {
	case "jpg", "jpeg":
		return filepath.Ext(path) == ".jpg" || filepath.Ext(path) == ".jpeg"
	default:
		return filepath.Ext(path) == "."+ext
	}
}

func genOutPath(path string, outExt string) (outPath string) {
	// return replaceFileExtent(path, filepath.Ext(path), "."+outExt)
	return replaceExt(path, outExt)
}

func replaceFileExtent(filePath string, oldExt, newExt string) string {
	if strings.HasSuffix(filePath, ".jpg") && oldExt == ".jpeg" {
		return replaceSuffix(filePath, ".jpg", string(newExt))
	} else if strings.HasSuffix(filePath, ".jpeg") && oldExt == ".jpg" {
		return replaceSuffix(filePath, ".jpeg", string(newExt))
	}
	return replaceSuffix(filePath, string(oldExt), string(newExt))
}

func replaceSuffix(s, old, new string) string {
	return strings.TrimSuffix(s, old) + new
}

func replaceExt(path, ext string) string {
	return strings.TrimSuffix(path, filepath.Ext(path)) + "." + ext
}
