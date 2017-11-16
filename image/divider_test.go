package image

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"image"
)

func TestDivide(t *testing.T) {
	assert := assert.New(t)

	rect := image.Rectangle{Min: image.Point{0,0}, Max: image.Point{100,100}}

	assert.Equal(100, rect.Max.X)
}
