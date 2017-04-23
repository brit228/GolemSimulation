package leapfrog

import (
	"vec3"
	. "structs"
)

func LeapfrogStep(typ int, mol *Molecule, deltaT Real) {
	var n Int = 0

	switch typ {
	case 1:
		for ; n<mol.NumAtoms; n++ {
			vec3.RAdd(mol.Atoms[n].Rv,vec3.RScale(mol.Atoms[n].Ra,0.5*deltaT))
			vec3.RAdd(mol.Atoms[n].R,vec3.RScale(mol.Atoms[n].Rv,deltaT))
		}
	case 2:
		for ; n<mol.NumAtoms; n++ {
			vec3.RAdd(mol.Atoms[n].Rv,vec3.RScale(mol.Atoms[n].Ra,0.5*deltaT))
		}
	}
}