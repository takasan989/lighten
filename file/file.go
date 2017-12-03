package file

import (
	"os"
	"fmt"
	"path/filepath"
)

// ファイルリストを絶対パスに変換する
func CheckFiles(files []string) ([]string, error) {
	files, err := abs(files)

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		stat, err := os.Stat(file)

		if err != nil || stat.IsDir() {
			return nil, fmt.Errorf("%s is not exists.", file)
		}
	}

	return files, nil
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
