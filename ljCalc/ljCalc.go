package ljCalc

import (
	. "GoLEM/structs"
	"GoLEM/vec3"
	"GoLEM/genfuncs"
	"math"
)

func Compute(rCut Real, virSum, vdwSum *Real, region VecR, mol *Molecule) {
	var rrCut Real
	var n Int

	rrCut = rCut * rCut
	for n=0; n<mol.NumAtoms; n++ {
		for n2 := range mol.Atoms[n].NebrList {
			EachCompute(mol, rrCut, vdwSum, virSum, n, Int(n2), region)
		}
	}
}

func EachCompute(mol *Molecule, rrCut Real, vdwSum, virSum *Real, n, n2 Int, region VecR) {
	var r, rr, eps, Rmin, rri, rri3, fcVal, uVal Real
	var dr VecR

	dr = vec3.RSub(mol.Atoms[n].R, mol.Atoms[n2].R,)
	dr = vec3.RWrapAll(dr, region)
	r = vec3.RLen(dr)
	rr = r * r
	if rr < rrCut {
		if !genfuncs.IntinExclusion(Int(n2), mol.Atoms[n].Exclusions) {
			eps = Real(math.Sqrt(float64(mol.Atoms[n].Eps * mol.Atoms[n2].Eps)))
			Rmin = mol.Atoms[n].Rmin * mol.Atoms[n2].Rmin
		} else {
			eps = Real(math.Sqrt(float64(mol.Atoms[n].Eps14 * mol.Atoms[n2].Eps14)))
			Rmin = mol.Atoms[n].Rmin14 * mol.Atoms[n2].Rmin14
		}
		rri = Rmin / rr
		rri3 = rri * rri * rri
		fcVal = 48. * eps * rri3 * (rri3 - 0.5) * rri
		uVal = 4. * eps * rri3 * (rri3 - 1.) + eps
		mol.Atoms[n].Ra = vec3.RAdd(mol.Atoms[n].Ra, vec3.RScale(dr, fcVal))
		mol.Atoms[n2].Ra = vec3.RAdd(mol.Atoms[n2].Ra, vec3.RScale(dr, -fcVal))
		*vdwSum += uVal
		*virSum += fcVal * rr
	}
}