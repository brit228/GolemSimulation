package main

import (
	leap "GolemSimulation/integrate/leapfrog"
	"GolemSimulation/boundary"
	nebr "GolemSimulation/speedup/ljspeedup/nebrlist"
)

func RunJob() {
	for ; stepCount<numSteps; {
		stepCount++
		leap.LeapfrogStep(1, &mol, deltaT)
		boundary.ApplyBoundaryConditions(&mol, region, wrapX, wrapY, wrapZ)
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