package main

import (
	. "GolemSimulation/structs"
	"GolemSimulation/vec3"
	"GolemSimulation/calc/ljCalc"
	"GolemSimulation/calc/electroCalc"
	"GolemSimulation/calc/angleCalc"
	"GolemSimulation/calc/bondCalc"
	"GolemSimulation/calc/dihedCalc"
	"GolemSimulation/calc/impropCalc"
	"GolemSimulation/calc/externalCalc"
)

func ComputeForces() {
	var n Int
	for n=0; n<mol.NumAtoms; n++ {
		mol.Atoms[n].Ra = vec3.RZero()
	}
	ljCalc.Compute(rCut, virSum, vdwSum, region, &mol)
	electroCalc.Compute()
	angleCalc.Compute()
	bondCalc.Compute()
	dihedCalc.Compute()
	impropCalc.Compute()
	externalCalc.Compute()
}