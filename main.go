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
	skipdirs := flag.String("skipdirs", "", "csv of dir names to skip")
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
	count, files, min, max := GetLines(d, strings.Split(e, ","), strings.Split(*skipdirs, ","))
	if files > 0 {
		avg := float64(count) / float64(files)
		avg = math.Round(avg)
		fmt.Println(d, "extension", e)
		fmt.Println("total", count, "files", files)
		fmt.Println("min", min, "max", max, "average", avg)
	}
	fmt.Println("Finished process in", time.Since(start))
}
