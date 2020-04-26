package main

import (
	"flag"

	"github.com/devlog/pkg"
)

func main() {
	outputFlagPtr := flag.String("p", "", "Path to the directory where devlog will save notes to. Only pass in a path to a directory, file names will be auto-generated.")
	templateFlagPtr := flag.String("c", "", "File path to the .yaml config file for your devlog notes.")
	flag.Parse()
	pkg.Start(*templateFlagPtr, *outputFlagPtr)
}
