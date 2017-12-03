package composite

import (
	"image"
	"fmt"
)

//
type CompositableSet interface {
	next() Compositable
}

type FileSet struct {
	files []string
}

func (set *FileSet) next() Compositable {
	if len(set.files) == 0 {
		return nil
	}

	file := set.files[0]
	set.files = set.files[1:]

	return &FileCompositable{file}
}

//
func Process(files []string) (image.Image, error) {
	compositor := ThroughCompositor{}
	set := &FileSet{files}
	compositable := set.next()

	if compositable == nil {
		return nil, fmt.Errorf("invalid\n")
	}

	next := set.next()
	for next != nil {
		compositable = compositor.composite(compositable, next)
		next = set.next()
	}

	return compositable.image()
}
