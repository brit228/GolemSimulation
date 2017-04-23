package vec3

import (
	. "GoLEM/structs"
	"math"
	"math/rand"
)

func RtoI(v VecR) VecI {
	var p VecI
	p.X = Int(v.X)
	p.Y = Int(v.Y)
	p.Z = Int(v.Z)
	return p
}

func ItoR(v VecI) VecR {
	var p VecR
	p.X = Real(v.X)
	p.Y = Real(v.Y)
	p.Z = Real(v.Z)
	return p
}

func RAdd (v1, v2 VecR) VecR {
	var v VecR
	v.X = v1.X + v2.X
	v.Y = v1.Y + v2.Y
	v.Z = v1.Z + v2.Z
	return v
}

func RSub (v1, v2 VecR) VecR {
	var v VecR
	v.X = v1.X - v2.X
	v.Y = v1.Y - v2.Y
	v.Z = v1.Z - v2.Z
	return v
}

func RMul (v1, v2 VecR) VecR {
	var v VecR
	v.X = v1.X * v2.X
	v.Y = v1.Y * v2.Y
	v.Z = v1.Z * v2.Z
	return v
}

func RDiv (v1, v2 VecR) VecR {
	var v VecR
	v.X = v1.X / v2.X
	v.Y = v1.Y / v2.Y
	v.Z = v1.Z / v2.Z
	return v
}

func RScale (v1 VecR, s Real) VecR {
	var v VecR
	v.X = v1.X * s
	v.Y = v1.Y * s
	v.Z = v1.Z * s
	return v
}

func RDot (v1, v2 VecR) Real {
	return v1.X * v2.X + v1.Y * v2.Y + v1.Z * v2.Z
}

func RLen(v VecR) Real {
	return Real(math.Sqrt(float64(RDot(v,v))))
}

func RProd(v VecR) Real {
	return v.X * v.Y * v.Z
}

func RVSum (v VecR) Real {
	return v.X + v.Y + v.Z
}

func RRand(randSource rand.Source) VecR {
	var scale float64 = 0.4656612873e-9
	var X, Y, s Real
	var v VecR

	s = 2.
	for ; s>1.; {
		X = 2.*Real(scale)*Real(randSource.Int63()) - 1
		Y = 2.*Real(scale)*Real(randSource.Int63()) - 1
		s = X * X + Y * Y
	}
	v.Z = 1. - 2. * s
	s = 2. * Real(math.Sqrt(1. - float64(s)))
	v.X = s * X
	v.Y = s * Y
	return v
}

func RWrapAll(v1, region VecR) VecR {
	var v VecR = v1
	if v.X >= 0.5*region.X {
		v.X -= region.X
	} else if v.X < -0.5*region.X {
		v.X += region.X
	}
	if v.Y >= 0.5*region.Y {
		v.Y -= region.Y
	} else if v.Y < -0.5*region.Y {
		v.Y += region.Y
	}
	if v.Z >= 0.5*region.Z {
		v.Z -= region.Z
	} else if v.Z < -0.5*region.Z {
		v.Z += region.Z
	}
	return v
}

func RZero() VecR {
	return VecR{0.,0.,0.}
}

func IAdd (v1, v2 VecI) VecI {
	var v VecI
	v.X = v1.X + v2.X
	v.Y = v1.Y + v2.Y
	v.Z = v1.Z + v2.Z
	return v
}

func ISub (v1, v2 VecI) VecI {
	var v VecI
	v.X = v1.X - v2.X
	v.Y = v1.Y - v2.Y
	v.Z = v1.Z - v2.Z
	return v
}

func IMul (v1, v2 VecI) VecI {
	var v VecI
	v.X = v1.X * v2.X
	v.Y = v1.Y * v2.Y
	v.Z = v1.Z * v2.Z
	return v
}

func IDiv (v1, v2 VecI) VecI {
	var v VecI
	v.X = v1.X / v2.X
	v.Y = v1.Y / v2.Y
	v.Z = v1.Z / v2.Z
	return v
}

func IScale (v1 VecI, s Int) VecI {
	var v VecI
	v.X = v1.X * s
	v.Y = v1.Y * s
	v.Z = v1.Z * s
	return v
}

func ILinear(v1, v2 VecI) Int {
	return (v1.Z * v2.Y + v1.Y) * v2.X + v1.X
}

func CellWrapAll(i1, i2 VecI, r1 VecR) (VecI, VecR) {
	var i VecI
	var r VecR

	if i1.X >= i2.X {
		i.X = 0
		r.X = r1.X
	} else if i1.X < 0 {
		i.X = i2.X - 1
		r.X = -r1.X
	}
	if i1.Y >= i2.Y {
		i.Y = 0
		r.Y = r1.Y
	} else if i1.Y < 0 {
		i.Y = i2.Y - 1
		r.Y = -r1.Y
	}
	if i1.Z >= i2.Z {
		i.Z = 0
		r.Z = r1.Z
	} else if i1.Z < 0 {
		i.Z = i2.Z - 1
		r.Z = -r1.Z
	}
	return i, r
}