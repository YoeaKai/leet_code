package simplify_path

import (
	"path/filepath"
	"strings"
)

// Method 1
func SimplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	clearDirs := []string{}

	for _, dir := range dirs {
		if dir == "." || dir == "/" || dir == "" {
			continue
		} else if dir == ".." {
			if len(clearDirs) > 0 {
				clearDirs = clearDirs[:len(clearDirs)-1]
			}
		} else {
			clearDirs = append(clearDirs, dir)
		}
	}

	return "/" + strings.Join(clearDirs, "/")
}

// Method 2
func SimplifyPath2(path string) string {
	return filepath.Clean(path)
}
