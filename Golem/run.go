package main

import (
	leap "leapfrog"
	"boundary"
	nebr "nebrlist"
)

func RunJob() {
	for ; stepCount<numSteps; {
		stepCount++
		leap.LeapfrogStep(1, &mol, deltaT)
		boundary.ApplyBoundaryConditions(&mol, region)
		nebr.NebrBuild()
		ComputeForces()
		leap.LeapfrogStep(2, &mol, deltaT)
		EvalProps()
		AccumProps(1)
		OutputPrinter()
		ApplyBarostat()
		ApplyThermostat()
		Analysis()
	}
}