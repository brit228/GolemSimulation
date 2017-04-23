package main

import (
	"os"
	"bufio"
	"strings"
	charmm "GoLEM/charmmInput"
	Gerror "GoLEM/errorCodes"
)

func ParseInputConfig(input string) {
	inputFile := OpenConfigFile(input)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		text := scanner.Text()
		splitText := strings.Fields(text)
		AllocateVar(splitText[0],splitText[1:])
	}
	InputRequirements()
}

func OpenConfigFile(file string) *os.File {
	f,err := os.Open(file)
	if err != nil {
		Gerror.GolemKill()Kill(2,"GoLEM: Invalid File Location or File!\n")
	}
	defer f.Close()

	return f
}

func AllocateVar(variable string, inputs []string) {
	if variable[0] != '#' {
		switch variable {
		case "charmmInput":
			if len(inputs) != 1 {
				Gerror.GolemKill(2, "GoLEM: Incorrect number of inputs for charmmInput parameter!\n")
			}
			charmm.CharmmInput = bool(inputs[1])
		case "psfInput":
			if len(inputs) != 1 {
				Gerror.GolemKill(2, "GoLEM: Incorrect number of inputs for psfInput parameter!\n")
			}
			charmm.PSFInput = inputs[1]
		}
	}
}

func InputRequirements() {

}