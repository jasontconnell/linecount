package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var exts string
var dir string

func main() {
	dir := flag.String("d", ".", "Starting directory")
	exts := flag.String("e", "", "Extensions. Extensions of files to count line numbers in")
	flag.Parse()

	if *dir == "" || *exts == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	d := *dir
	e := *exts

	if !filepath.IsAbs(d) {
		wd, _ := os.Getwd()
		if d != "." {
			d = filepath.Join(wd, d)
		} else {
			d = wd
		}
	}

	start := time.Now()
	count, files := GetLines(d, strings.Split(e, ","))
	if files > 0 {
		avg := float64(count) / float64(files)
		avg = math.Round(avg)
		fmt.Println(d, "has", count, "LOC in", files, "files with extensions in", e, ". Average", int64(avg))
	}
	fmt.Println("Finished process in", time.Since(start))
}
