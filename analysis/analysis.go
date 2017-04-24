package analysis

import (
	"github.com/robertkrimen/otto"
	"os"
	"io/ioutil"
	Gerror "GolemSimulation/errorCodes"
	. "GolemSimulation/structs"
)

var analysisVM *otto.Otto
var analysisBool bool
var analysisSetupBool bool
var analysisFile string
var analysisSetupFile string

func AnalysisSetup() {
	if analysisSetupBool {
		AnalysisReadFile(analysisSetupFile)
		AnalysisReadFile(analysisFile)
		analysisVM = otto.New()
		analysisVM.Run(analysisSetupFile)
	} else if analysisBool {
		AnalysisReadFile(analysisFile)
		analysisVM = otto.New()
	}
}

func Analysis() {
	if analysisBool {
		analysisVM.Run(analysisFile)
	}
}

func AnalysisReadFile(file string) string {
	f,err := os.Open(file)
	if err != nil {
		Gerror.GolemKill(2,"Golem: Error Opening Analysis Script!\n")
	}
	defer f.Close()

	byteFile,err := ioutil.ReadFile(file)
	fileString := string(byteFile)

	return fileString
}

func AnalysisJSFuncs(mol *Molecule) {
	analysisVM.Set("atomPos", func(call otto.FunctionCall) otto.Value {
		id,_ := call.Argument(0).ToInteger()
		result,_ := analysisVM.ToValue(mol.Atoms[id].R)
		return result
	})
	analysisVM.Set("atomVel", func(call otto.FunctionCall) otto.Value {
		id,_ := call.Argument(0).ToInteger()
		result,_ := analysisVM.ToValue(mol.Atoms[id].Rv)
		return result
	})
}