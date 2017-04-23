package main

func main() {
	ParseCommandLine()
	ParseInputConfig(inputConfigFile)
	SetupJob()
	RunJob()
}