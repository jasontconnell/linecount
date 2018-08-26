package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func GetLines(dir string, exts []string) (count, files int64) {
	extmap := make(map[string]string)
	for _, ext := range exts {
		key := "." + strings.ToLower(ext)
		extmap[key] = key
	}

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		ext := strings.ToLower(filepath.Ext(path))
		if _, ok := extmap[ext]; ok {
			files++
			f, err := os.Open(path)
			if err != nil {
				return err
			}

			s := bufio.NewScanner(f)
			for s.Scan() {
				count++
			}
		}

		return nil
	})

	return
}
