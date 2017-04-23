package main

import (
	"flag"
	"runtime"
	Gerror "GoLEM/errorCodes"
)

func ParseCommandLine() {
	flag.StringVar(&inputConfigFile, "i", "", "Input file for GoMM")
	flag.IntVar(&inputNumThreads, "p", -1, "Max number of processors for GoMM")
	flag.Parse()
	if inputConfigFile == "" {
		Gerror.GolemKill(2,"GoMM: No Input File Found!\n")
	}
	runtime.GOMAXPROCS(inputNumThreads)
}