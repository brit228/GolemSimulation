package main

import (
	. "GolemSimulation/structs"
	"GolemSimulation/vec3"
	lj "GolemSimulation/calc/ljCalc"
	electro "GolemSimulation/calc/electroCalc"
)

func ComputeForces() {
	var n Int
	for n=0; n<mol.NumAtoms; n++ {
		mol.Atoms[n].Ra = vec3.RZero()
	}
	lj.Compute(rCut, virSum, vdwSum, region, &mol)
	electro.Compute()
}