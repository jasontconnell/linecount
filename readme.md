*Usage*

linecount -d . -e go

>. has 69 LOC in 2 files with extensions in go . Average 34.5

> Finished process in 1.9769ms

API included

GetLines(dir string, exts []string) (lines, files int64)

lines, files := GetLines(".", []string{"go"})

>lines = 69, files = 2

TODO: Process multiple files concurrently