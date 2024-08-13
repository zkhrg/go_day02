package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zkhrg/go_day02/pkg/counter"
)

func main() {
	wordsFlag := flag.Bool("w", false, "count words")
	linesFlag := flag.Bool("l", false, "count lines")
	symbolsFlag := flag.Bool("m", false, "count symbols")

	opt := counter.Opt{
		Symb: false,
		Word: false,
		Line: false,
	}

	flag.Parse()

	if *wordsFlag && !(*linesFlag) && !(*symbolsFlag) {
		opt.Word = true
	} else if !(*wordsFlag) && *linesFlag && !(*symbolsFlag) {
		opt.Line = true
	} else if !(*wordsFlag) && !(*linesFlag) && *symbolsFlag {
		opt.Symb = true
	} else {
		fmt.Fprintf(os.Stderr, "usage: <program> <flags: -w | -l | -m | (reqiered ONE)> <files...>")
		os.Exit(1)
	}
	if len(flag.Args()) == 0 {
		fmt.Fprintf(os.Stderr, "usage: <program> <flags: -w | -l | -m | (reqiered ONE)> <files...>")
		os.Exit(1)
	}
	args := flag.Args()
	counter.Process(opt, args)
}
