package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zkhrg/go_day02/pkg/walker"
)

func main() {
	symlinkFlag := flag.Bool("ls", false, "output symlinks")
	dirFlag := flag.Bool("d", false, "output dirs")
	fileFlag := flag.Bool("f", false, "output files")
	extFlag := flag.String("ext", "", "accept file extension")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 || (*extFlag != "" && !(*fileFlag)) {
		fmt.Println("Usage: <program> <flags: -ls | -d | -f | -ext (reqiered f)> <directory>")
		os.Exit(1)
	}

	directory := args[0]

	var opt walker.Opt
	if !(*symlinkFlag) && !(*dirFlag) && !(*fileFlag) {
		opt = walker.Opt{
			Dir:     true,
			File:    true,
			Symlink: true,
			Ext:     "",
		}
	} else {
		opt = walker.Opt{
			Dir:     *dirFlag,
			File:    *fileFlag,
			Symlink: *symlinkFlag,
			Ext:     *extFlag,
		}
	}

	walker.Find(directory, &opt)
}
