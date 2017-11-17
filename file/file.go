package file

import (
	"os"
	"fmt"
	"path/filepath"
)

func CheckFiles(files []string) (bool, error) {
	files, err := abs(files)

	if err != nil {
		return false, err
	}

	for _, file := range files {
		stat, err := os.Stat(file)

		if err != nil || stat.IsDir() {
			return false, fmt.Errorf("%s is not exists.", file)
		}
	}

	return true, nil
}

func abs(files[] string) ([]string, error) {
	abs := []string{}
	for _, file := range files {
		path, err := filepath.Abs(file)
		if err != nil {
			return nil, err
		}
		abs = append(abs, path)
	}
	return abs, nil
}
