package main

import (
	"fmt"
	"flag"
	f "github.com/takasan989/lighten/file"
	"os"
)

type Option struct {
	output string
}

var opt Option

func main() {
	flag.StringVar(&opt.output, "output", "", "output file")
	flag.Parse()

	_, err := f.CheckFiles(flag.Args())
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
