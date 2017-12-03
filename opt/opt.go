package opt

import (
	"flag"
)

type Option struct {
	Output string
	Files []string
}

func Parse() *Option {
	opt := &Option{}

	flag.StringVar(&opt.Output, "output", "", "output file")

	flag.Parse()
	opt.Files = flag.Args()

	return opt
}
