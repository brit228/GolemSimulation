package structs

type Real float64
type Int int64

type VecR struct {
	X, Y, Z Real
}
type VecI struct {
	X, Y, Z Int
}

type Atom struct {
	R, Rv, Ra VecR
	Mass, Charge, Eps, Rmin, Eps14, Rmin14 Real
	Typ1, Typ2 string
	ID Int
	NebrList []Int
	Exclusions []Exclusion
}

type Bond struct {
	Atom1, Atom2 Int
	Kb, B0 Real
}

type Angle struct {
	Atom1, Atom2, Atom3 Int
	Ktheta, Ktheta0 Real
}

type Dihedral struct {
	Atom1, Atom2, Atom3, Atom4 Int
	N Int
	Kchi, Delta Real
}

type Improper struct {
	Atom1, Atom2, Atom3, Atom4 Int
	Kpsi, Psi0 Real
}

type Exclusion struct {
	Atom2 Int
	Scaling Real
}

type Molecule struct {
	Atoms []Atom
	Bonds []Bond
	Angles []Angle
	Dihedrals []Dihedral
	Impropers []Improper

	NumAtoms Int
	NumBonds Int
	NumDihedrals Int
	NumImpropers Int
}

type Prop struct {
	Val, Sum, Sum2 Real
}