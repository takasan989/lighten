package main

import (
	"fmt"
	"flag"
	"github.com/takasan989/lighten/opt"
	f "github.com/takasan989/lighten/file"
	i "github.com/takasan989/lighten/image"
	"os"
	"image"
	"io"
	"image/jpeg"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
}

func main() {
	opt := opt.Parse()
	files, err := f.CheckFiles(opt.Files)

	check(err)

	image, err := i.Process(files)

	check(err)

	if opt.Output == "" {
		writer := os.Stdout
		write(image, writer)
	} else {
		abs, err := filepath.Abs(opt.output)

		check(err)

		fp, err := os.Create(abs)

		check(err)

		defer fp.Close()

		write(image, fp)
	}
}

func write(image image.Image, writer io.Writer) error {
	return jpeg.Encode(writer, image, &jpeg.Options{Quality: 100})
}
