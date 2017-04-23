package nebrlist

import (
	. "GolemSimulation/structs"
	"GolemSimulation/vec3"
)

func NebrBuild(rCut, rNebrShell Real, nebrTableLen, nOffset Int, cells VecI, region VecR, mol *Molecule, cellList []Int, vOff []VecI) {
	var rrNebr Real
	var c, n, m1x, m1y, m1z, m1, offset Int
	var invWid, rs, shift VecR
	var cc, m1v, m2v VecI

	rrNebr = (rCut + rNebrShell) * (rCut + rNebrShell)
	invWid = vec3.RDiv(vec3.ItoR(cells), region)

	for n=0; n<mol.NumAtoms; n++ {
		rs = vec3.RAdd(mol.Atoms[n].R, vec3.RScale(region, 0.5))
		cc = vec3.RtoI(vec3.RMul(rs, invWid))
		c = vec3.ILinear(cc, cells) + mol.NumAtoms
		cellList[n] = cellList[c]
		cellList[c] = n
	}

	nebrTableLen = 0
	for m1z=0; m1z<cells.Z; m1z++ {
		for m1y=0; m1y<cells.Y; m1y++ {
			for m1x=0; m1x<cells.X; m1x++ {
				m1v = VecI{m1x, m1y, m1z}
				m1 = vec3.ILinear(m1v, cells) + mol.NumAtoms
				for offset=0; offset<nOffset; offset++ {
					m2v = vec3.IAdd(m1v, vOff[offset])
					shift = vec3.RZero()

				}
			}
		}
	}
}