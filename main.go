package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var exts string
var dir string

func init() {
	flag.StringVar(&dir, "d", "", "Starting directory")
	flag.StringVar(&exts, "e", "", "Extensions. Extensions of files to count line numbers in")
}

func main() {
	flag.Parse()

	if dir == "" || exts == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	start := time.Now()
	count, files := GetLines(dir, strings.Split(exts, ","))
	if files > 0 {
		fmt.Println(dir, "has", count, "LOC in", files, "files with extensions in", exts, ". Average", float64(count)/float64(files))
	}
	fmt.Println("Finished process in", time.Since(start))
}
