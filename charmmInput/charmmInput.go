package charmmInput

import (
	. "structs"
	"os"
	"bufio"
	"strings"
	"strconv"
	Gerror "errorCodes"
)

var CharmmInput bool
var PSFInput string

func ParsePSF(psfinput string, mol *Molecule) {
	var nMax int64
	var typeCase string

	fileOpen,err := os.Open(psfinput)
	if err != nil {
		Gerror.GolemKill(2,"GoMM: Invalid PSF file or file location!\n")
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
					Gerror.GolemKill(2,"GoMM: Invalid PSF file formatting!\n")
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
					Gerror.GolemKill(2,"GoMM: Invalid PSF file formatting!\n")
				}
				atom.ID = Int(id - 1)
				mol.Atoms[atom.ID] = atom
			}
		}
	}
	if err = scanner.Err(); err != nil {
		Gerror.GolemKill(2,"GoMM: Invalid PSF file formatting!\n")
	}
}