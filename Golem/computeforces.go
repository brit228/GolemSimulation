package main

import (
	lj "ljCalc"
)

func ComputeForces() {
	lj.Compute(rCut, virSum, vdwSum, region, mol)
}