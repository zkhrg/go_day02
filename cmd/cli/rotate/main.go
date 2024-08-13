package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/zkhrg/go_day02/pkg/rotater"
)

func main() {
	archiveFlag := flag.String("a", ".", "specify archive folder")

	flag.Parse()

	args := flag.Args()
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	err := os.MkdirAll(*archiveFlag, os.ModePerm)
	if err != nil || len(args) == 0 {
		fmt.Fprintf(os.Stderr, "usage: <program> <flags: -a (not reqiered)> <files...>")
		os.Exit(1)
	}

	rotater.LogRotate(*archiveFlag, timestamp, args)
}
