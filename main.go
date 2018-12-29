package main

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"sync"
)

const (
	dosEOL = "\r\n"
	nixEOL = "\n"
)

func main() {
	var (
		err    error
		params = getOptParams()
	)

	getDebugLogger(params.Build.Debug)

	eol := nixEOL
	if params.Runtime.AddCarriage {
		eol = dosEOL
	}

	files := []string{
		params.Runtime.File,
	}

	if len(params.Runtime.FilePattern) > 0 {
		files, err = filepath.Glob(params.Runtime.FilePattern)
		if err != nil {
			logger.Fatal(err)
		}
	}

	var wait sync.WaitGroup

	for _, file := range files {
		source := openFile(file)
		wait.Add(1)
		go func() {
			defer source.Close()
			defer wait.Done()
			dest := readFile(source, eol)

			if params.Runtime.Inplace {
				writeFile(source, dest)
			} else {
				os.Stdout.WriteString(dest.String())
			}
		}()
	}
	wait.Wait()
}

func openFile(fname string) *os.File {
	stats, err := os.Stat(fname)
	if err != nil {
		logger.Println(err)
	}

	source, err := os.OpenFile(fname, os.O_RDWR, stats.Mode())
	if err != nil {
		logger.Println(err)
	}
	return source
}

func readFile(source *os.File, eol string) bytes.Buffer {
	var dest bytes.Buffer
	scanner := bufio.NewScanner(source)

	for scanner.Scan() {
		dest.Write(scanner.Bytes())
		dest.WriteString(eol)
	}

	if err := scanner.Err(); err != nil {
		logger.Println(err)
	}
	return dest
}

func writeFile(source *os.File, dest bytes.Buffer) {
	source.Truncate(0)
	source.Seek(0, os.SEEK_SET)

	n, err := source.Write(dest.Bytes())
	if err != nil {
		logger.Println(err)
	}

	logger.Println("bytes written:", n)
}
