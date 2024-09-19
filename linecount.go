package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

func GetLines(dir string, exts []string, skipdirs []string) (count, files, min, max int64) {
	min = math.MaxInt32
	max = math.MinInt32

	extmap := make(map[string]string)
	for _, ext := range exts {
		key := "." + strings.ToLower(ext)
		extmap[key] = key
	}

	dirmap := make(map[string]string)
	for _, skip := range skipdirs {
		dirmap[skip] = skip
	}

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			_, ok := dirmap[info.Name()]
			if ok || strings.HasPrefix(info.Name(), ".") {
				log.Println("skipping", path)
				return filepath.SkipDir
			}
		}
		ext := strings.ToLower(filepath.Ext(path))
		if _, ok := extmap[ext]; ok {
			files++
			f, err := os.Open(path)
			if err != nil {
				return err
			}

			var lines int64
			s := bufio.NewScanner(f)
			for s.Scan() {
				lines++
			}
			if lines < min {
				min = lines
			}
			if lines > max {
				max = lines
			}
			count += lines
		}

		return nil
	})

	return
}
