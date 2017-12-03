package composite

import (
	"image"
	"fmt"
)

type Compositor interface {
	composite(Compositable, Compositable) (Compositable, error)
}

type ThroughCompositor struct {}

func (compositor *ThroughCompositor) composite(first, second Compositable) (Compositable, error) {
	return second, nil
}

type BrightnessCompositor struct {}

func (_ *BrightnessCompositor) composite(first, second Compositable) (Compositable, error) {
	fstImage, err := first.image()
	if err != nil {
		return nil, err
	}
	
	scdImage, err := second.image()
	if err != nil {
		return nil, err
	}

	if checkBounds(fstImage, scdImage) {
		return nil, fmt.Errorf("size is must be same")
	}


}

func checkBounds(a, b image.Image) bool {
	return a.Bounds().Eq(b.Bounds())
}
