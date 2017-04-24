package psfInput

import (
	. "GolemSimulation/structs"
	"os"
	"bufio"
	"strings"
	"strconv"
	Gerror "GolemSimulation/errorCodes"
)

var atomNum int = 0
var bondNum int = 0
var angleNum int = 0
var dihedNum int = 0
var impropNum int = 0

func Parse(psfinput string, mol *Molecule) {
	var nMax int64
	var typeCase string

	fileOpen,err := os.Open(psfinput)
	if err != nil {
		Gerror.GolemKill(2,"Golem: Invalid PSF file or file location!\n")
	}
	defer fileOpen.Close()

	scanner := bufio.NewScanner(fileOpen)
	for scanner.Scan() {
		text := scanner.Text()
		splitText := strings.Fields(text)
		if splitText[0] != "REMARKS" {
			if splitText[1][0] == '!' {
				nMax, err = strconv.ParseInt(splitText[0], 10, 64)
				if err != nil {
					Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
				}
				switch splitText[1] {
				case "!NATOM":
					mol.Atoms = make([]Atom, nMax)
					typeCase = "ATOM"
				case "!NBOND":
					mol.Bonds = make([]Bond, nMax)
					typeCase = "BOND"
				case "!NTHETA":
					mol.Bonds = make([]Bond, nMax)
					typeCase = "ANGLE"
				case "!NPHI":
					mol.Bonds = make([]Bond, nMax)
					typeCase = "DIHEDRAL"
				case "!NIMPHI":
					mol.Bonds = make([]Bond, nMax)
					typeCase = "IMPROPER"
				}
			}
			switch typeCase {
			case "ATOM":
				var atom Atom

				id,err := strconv.ParseInt(text[0:10], 10, 64)
				if err != nil {
					Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
				}
				atom.ID = Int(id - 1)
				mol.Atoms[atom.ID] = atom
				atomNum++
			case "BOND":
				var bond Bond

				if !(bool(len(splitText) % 2)) {
					Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
				}
				for n:=0; n<len(splitText)/2; n+=2 {
					if n >= 8 {
						Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
					}
					id1,err := strconv.ParseInt(splitText[n], 10, 64)
					if err != nil {
						Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
					}
					id2,err := strconv.ParseInt(splitText[n+1], 10, 64)
					if err != nil {
						Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
					}
					bond.Atom1 = Int(id1)
					bond.Atom2 = Int(id2)
					mol.Bonds[bondNum] = bond
					bondNum++
				}
			case "ANGLE":
				var angle Angle

				if !(bool(len(splitText) % 3)) {
					Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
				}
				for n:=0; n<len(splitText)/3; n+=3 {
					if n >= 6 {
						Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
					}
					id1,err := strconv.ParseInt(splitText[n], 10, 64)
					if err != nil {
						Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
					}
					id2,err := strconv.ParseInt(splitText[n+1], 10, 64)
					if err != nil {
						Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
					}
					id3,err := strconv.ParseInt(splitText[n+2], 10, 64)
					if err != nil {
						Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
					}
					angle.Atom1 = Int(id1)
					angle.Atom2 = Int(id2)
					angle.Atom3 = Int(id3)
					mol.Angles[angleNum] = angle
					angleNum++
				}
			case "DIHEDRAL":
				var dihed Dihedral

				if !(bool(len(splitText) % 4)) {
					Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
				}
				for n:=0; n<len(splitText)/3; n+=3 {
					if n >= 8 {
						Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
					}
					id1,err := strconv.ParseInt(splitText[n], 10, 64)
					if err != nil {
						Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
					}
					id2,err := strconv.ParseInt(splitText[n+1], 10, 64)
					if err != nil {
						Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
					}
					id3,err := strconv.ParseInt(splitText[n+2], 10, 64)
					if err != nil {
						Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
					}
					id4,err := strconv.ParseInt(splitText[n+2], 10, 64)
					if err != nil {
						Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
					}
					dihed.Atom1 = Int(id1)
					dihed.Atom2 = Int(id2)
					dihed.Atom3 = Int(id3)
					dihed.Atom4 = Int(id4)
					mol.Dihedrals[dihedNum] = dihed
					dihedNum++
				}
			case "IMPROPER":
				var improp Improper

				if !(bool(len(splitText) % 4)) {
					Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
				}
				for n:=0; n<len(splitText)/3; n+=3 {
					if n >= 8 {
						Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
					}
					id1,err := strconv.ParseInt(splitText[n], 10, 64)
					if err != nil {
						Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
					}
					id2,err := strconv.ParseInt(splitText[n+1], 10, 64)
					if err != nil {
						Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
					}
					id3,err := strconv.ParseInt(splitText[n+2], 10, 64)
					if err != nil {
						Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
					}
					id4,err := strconv.ParseInt(splitText[n+2], 10, 64)
					if err != nil {
						Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
					}
					improp.Atom1 = Int(id1)
					improp.Atom2 = Int(id2)
					improp.Atom3 = Int(id3)
					improp.Atom4 = Int(id4)
					mol.Impropers[impropNum] = improp
					impropNum++
				}
			}
		}
	}
	if len(mol.Atoms) != atomNum {
		Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
	}
	if len(mol.Bonds) != bondNum {
		Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
	}
	if len(mol.Angles) != angleNum {
		Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
	}
	if len(mol.Dihedrals) != dihedNum {
		Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
	}
	if len(mol.Impropers) != impropNum {
		Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
	}
	if err = scanner.Err(); err != nil {
		Gerror.GolemKill(2,"Golem: Invalid PSF file formatting!\n")
	}
}