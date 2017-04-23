package main

import (
	. "GoLEM/structs"
	"GoLEM/vec3"
	lj "GoLEM/ljCalc"
	electro "GoLEM/electroCalc"
)

func ComputeForces() {
	var n Int
	for n=0; n<mol.NumAtoms; n++ {
		mol.Atoms[n].Ra = vec3.RZero()
	}
	lj.Compute(rCut, virSum, vdwSum, region, &mol)
	electro.Compute()
}