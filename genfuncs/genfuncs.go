package genfuncs

import (
	. "structs"
)

func IntinExclusion(val Int, exclusions []Exclusion) bool {
	var dat Exclusion
	for dat = range exclusions {
		if dat.Atom2 == val {
			return true
		}
	}
	return false
}