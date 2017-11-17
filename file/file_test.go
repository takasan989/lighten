package file

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestCheckFiles(t *testing.T) {
	assert := assert.New(t)

	dir, _ := os.Getwd()
	files := []string{
		path.Join(dir, "file.go"),
		path.Join(dir, "file_test.go"),
	}

	assert.True(CheckFiles(files))

	files = []string{
		path.Join(dir, "file.go"),
		path.Join(dir, "file_hoge.go"),
	}

	exists, err := CheckFiles(files)
	assert.False(exists)
	assert.NotNil(err)

	files = []string{
		path.Join(dir, "./"),
	}

	exists, err = CheckFiles(files)
	assert.False(exists)
	assert.NotNil(err)
}
