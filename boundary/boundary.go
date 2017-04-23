package boundary

import (
	. "GoLEM/structs"
	"GoLEM/vec3"
)

func ApplyBoundaryConditions(mol *Molecule, region VecR) {
	var n Int

	for ; n<mol.NumAtoms; n++ {
		mol.Atoms[n].R = vec3.RWrapAll(mol.Atoms[n].R, region)
	}
}