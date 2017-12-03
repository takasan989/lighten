package composite

import (
	"image"
	"image/jpeg"
	"os"
)

type Compositable interface {
	image() (image.Image, error)
}

type FileCompositable struct {
	path string
}

func (file *FileCompositable) image() (image.Image, error) {
	fp, err := os.Open(file.path)

	if err != nil {
		return nil, err
	}
	defer fp.Close()

	return jpeg.Decode(fp)
}

type Memory struct {
	img image.Image
}

func (memory *Memory) image() (image.Image, error) {
	return memory.img, nil
}
