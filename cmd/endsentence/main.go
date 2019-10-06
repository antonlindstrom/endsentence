package main

import (
	"flag"
	"fmt"

	"github.com/antonlindstrom/endsentence/endsentence"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var buildref string

func main() {
	version := flag.Bool("version", false, "Print build version")
	flag.Parse()

	if *version {
		fmt.Printf("endsentence build: %v\n", buildref)
		return
	}

	singlechecker.Main(endsentence.Analyzer)
}
