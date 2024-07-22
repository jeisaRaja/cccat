package main

import (
	"cccat/program"
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cccat <filepath>")
		os.Exit(1)
	}
	ln := flag.Bool("n", false, "show line number")
	flag.Parse()
	p := program.Program{LineBool: *ln}

	for _, arg := range flag.Args() {
		p.AddFilepaths(arg)
	}

	p.ReadFiles()
}
