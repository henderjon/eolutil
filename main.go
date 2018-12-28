package main

import (
	"bufio"
	"bytes"
	"os"
)

const (
	dosEOL = "\r\n"
	nixEOL = "\n"
)

func main() {
	params := getOptParams()
	getDebugLogger(params.Build.Debug)

	stats, err := os.Stat(params.Runtime.File)
	if err != nil {
		logger.Println(err)
	}

	source, err := os.OpenFile(params.Runtime.File, os.O_RDWR, stats.Mode())
	if err != nil {
		logger.Println(err)
	}
	defer source.Close()

	var dest bytes.Buffer
	scanner := bufio.NewScanner(source)

	eol := nixEOL
	if params.Runtime.AddCarriage {
		eol = dosEOL
	}

	for scanner.Scan() {
		dest.Write(scanner.Bytes())
		dest.WriteString(eol)
	}

	if err := scanner.Err(); err != nil {
		logger.Println(err)
	}

	if params.Runtime.Inplace {
		source.Truncate(0)
		source.Seek(0, os.SEEK_SET)

		n, err := source.Write(dest.Bytes())
		if err != nil {
			logger.Println(err)
		}

		logger.Println("bytes written:", n)
	} else {
		os.Stdout.WriteString(dest.String())
	}

}
