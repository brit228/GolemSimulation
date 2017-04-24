package boundary

import (
	. "GolemSimulation/structs"
	"GolemSimulation/vec3"
)

func ApplyBoundaryConditions(mol *Molecule, region VecR, wrapX, wrapY, wrapZ bool) {
	var n Int

	for ; n<mol.NumAtoms; n++ {
		if wrapX && wrapY && wrapZ {
			mol.Atoms[n].R = vec3.RWrapAll(mol.Atoms[n].R, region)
		} else {
			if wrapX {
				mol.Atoms[n].R = vec3.RWrapX(mol.Atoms[n].R, region)
			}
			if wrapY {
				mol.Atoms[n].R = vec3.RWrapY(mol.Atoms[n].R, region)
			}
			if wrapZ {
				mol.Atoms[n].R = vec3.RWrapZ(mol.Atoms[n].R, region)
			}
		}
	}
}