package main

import (
	"flag"
	"runtime"
	"fmt"
	Gerror "GolemSimulation/errorCodes"
)

func ParseCommandLine() {
	flag.StringVar(&inputConfigFile, "i", "", "Input file for Golem")
	flag.IntVar(&inputNumThreads, "p", -1, "Max number of processors for Golem")
	flag.BoolVar(&listLibrary, "list", false, "List packages used in Golem binary")
	flag.Parse()
	if listLibrary {
		fmt.Printf(library)
	} else {
		if inputConfigFile == "" {
			Gerror.GolemKill(2, "Golem: No Input File Found!\n")
		}
		if inputNumThreads <= 0 || inputNumThreads > runtime.NumCPU() {
			Gerror.GolemKill(2, "Golem: Incorrect Number of Processors!\n")
		}
		runtime.GOMAXPROCS(inputNumThreads)
	}
}