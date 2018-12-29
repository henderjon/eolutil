package main

import (
	"flag"
	"fmt"
	"os"
)

type buildParams struct {
	Timestamp string
	Version   string
	Debug     bool
}

type runtimeParams struct {
	File        string
	FilePattern string
	Inplace     bool
	AddCarriage bool
}

type getOptParameters struct {
	Build   buildParams
	Runtime runtimeParams
	Help    bool
}

func getOptParams() *getOptParameters {
	params := &getOptParameters{}
	flag.BoolVar(&params.Build.Debug, "debug", false, "once more, with feeling")
	flag.StringVar(&params.Runtime.File, "f", "", "the file on which to act")
	flag.StringVar(&params.Runtime.FilePattern, "fs", "", "glob pattern for multiple files; note that the pattern should be double quoted or escaped to prevent shell from globbing first")
	flag.BoolVar(&params.Runtime.AddCarriage, "rn", false, "use \\r\\n instead of \\n")
	flag.BoolVar(&params.Runtime.Inplace, "i", false, "modify the file in place instead of echoing to STDOUT")
	flag.BoolVar(&params.Help, "help", false, "show this message")
	flag.Parse()

	if params.Help {
		fmt.Println("built:", buildTimestamp)
		fmt.Println("version:", buildVersion)
		flag.PrintDefaults()
		os.Exit(0)
	}

	// value, ok := os.LookupEnv("")

	// set globally via linker during compilation; in version.go
	params.Build.Timestamp = getBuildTimestamp()
	params.Build.Version = getBuildVersion()

	return params
}
