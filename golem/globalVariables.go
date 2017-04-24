package main

import (
	. "GolemSimulation/structs"
)

var inputConfigFile string
var inputNumThreads int
var listLibrary bool

var stepCount Int
var numSteps Int
var deltaT Real

var region VecR
var mol Molecule

var wrapX, wrapY, wrapZ bool