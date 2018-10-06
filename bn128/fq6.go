package bn128

import (
	"math/big"
)

// Fq6 is Field 6
type Fq6 struct {
	F          Fq2
	NonResidue [2]*big.Int
}

// NewFq6 generates a new Fq6
func NewFq6(f Fq2, nonResidue [2]*big.Int) Fq6 {
	fq6 := Fq6{
		f,
		nonResidue,
	}
	return fq6
}

// Zero returns a Zero value on the Fq6
func (fq6 Fq6) Zero() [3][2]*big.Int {
	return [3][2]*big.Int{fq6.F.Zero(), fq6.F.Zero(), fq6.F.Zero()}
}

// One returns a One value on the Fq6
func (fq6 Fq6) One() [3][2]*big.Int {
	return [3][2]*big.Int{fq6.F.One(), fq6.F.One(), fq6.F.One()}
}

func (fq6 Fq6) mulByNonResidue(a [2]*big.Int) [2]*big.Int {
	return fq6.F.Mul(fq6.NonResidue, a)
}

// Add performs an addition on the Fq6
func (fq6 Fq6) Add(a, b [3][2]*big.Int) [3][2]*big.Int {
	return [3][2]*big.Int{
		fq6.F.Add(a[0], b[0]),
		fq6.F.Add(a[1], b[1]),
		fq6.F.Add(a[2], b[2]),
	}
}

// Sub performs a substraction on the Fq6
func (fq6 Fq6) Sub(a, b [3][2]*big.Int) [3][2]*big.Int {
	return [3][2]*big.Int{
		fq6.F.Sub(a[0], b[0]),
		fq6.F.Sub(a[1], b[1]),
		fq6.F.Sub(a[2], b[2]),
	}
}

// Neg performs a negation on the Fq6
func (fq6 Fq6) Neg(a [3][2]*big.Int) [3][2]*big.Int {
	return fq6.Sub(fq6.Zero(), a)
}

// Mul performs a multiplication on the Fq6
func (fq6 Fq6) Mul(a, b [3][2]*big.Int) [3][2]*big.Int {
	v0 := fq6.F.Mul(a[0], b[0])
	v1 := fq6.F.Mul(a[1], b[1])
	v2 := fq6.F.Mul(a[2], b[2])
	return [3][2]*big.Int{
		fq6.F.Add(
			v0,
			fq6.mulByNonResidue(
				fq6.F.Sub(
					fq6.F.Mul(
						fq6.F.Add(a[1], a[2]),
						fq6.F.Add(b[1], b[2])),
					fq6.F.Add(v1, v2)))),

		fq6.F.Add(
			fq6.F.Sub(
				fq6.F.Mul(
					fq6.F.Add(a[0], a[1]),
					fq6.F.Add(b[0], b[1])),
				fq6.F.Add(v0, v1)),
			fq6.mulByNonResidue(v2)),

		fq6.F.Add(
			fq6.F.Sub(
				fq6.F.Mul(
					fq6.F.Add(a[0], a[2]),
					fq6.F.Add(b[0], b[2])),
				fq6.F.Add(v0, v2)),
			v1),
	}
}

// Inverse returns the inverse on the Fq6
func (fq6 Fq6) Inverse(a [3][2]*big.Int) [3][2]*big.Int {
	t0 := fq6.F.Square(a[0])
	t1 := fq6.F.Square(a[1])
	t2 := fq6.F.Square(a[2])
	t3 := fq6.F.Mul(a[0], a[1])
	t4 := fq6.F.Mul(a[0], a[2])
	t5 := fq6.F.Mul(a[1], a[2])

	c0 := fq6.F.Sub(t0, fq6.mulByNonResidue(t5))
	c1 := fq6.F.Sub(fq6.mulByNonResidue(t2), t3)
	c2 := fq6.F.Sub(t1, t4)

	t6 := fq6.F.Inverse(
		fq6.F.Add(
			fq6.F.Mul(a[0], c0),
			fq6.mulByNonResidue(
				fq6.F.Add(
					fq6.F.Mul(a[2], c1),
					fq6.F.Mul(a[1], c2)))))
	return [3][2]*big.Int{
		fq6.F.Mul(t6, c0),
		fq6.F.Mul(t6, c1),
		fq6.F.Mul(t6, c2),
	}
}

// Div performs a division on the Fq6
func (fq6 Fq6) Div(a, b [3][2]*big.Int) [3][2]*big.Int {
	return fq6.Mul(a, fq6.Inverse(b))
}

// Square performs a square operation on the Fq6
func (fq6 Fq6) Square(a [3][2]*big.Int) [3][2]*big.Int {
	s0 := fq6.F.Square(a[0])
	ab := fq6.F.Mul(a[0], a[1])
	s1 := fq6.F.Add(ab, ab)
	s2 := fq6.F.Square(
		fq6.F.Add(
			fq6.F.Sub(a[0], a[1]),
			a[2]))
	bc := fq6.F.Mul(a[1], a[2])
	s3 := fq6.F.Add(bc, bc)
	s4 := fq6.F.Square(a[2])

	return [3][2]*big.Int{
		fq6.F.Add(
			s0,
			fq6.mulByNonResidue(s3)),
		fq6.F.Add(
			s1,
			fq6.mulByNonResidue(s4)),
		fq6.F.Sub(
			fq6.F.Add(
				fq6.F.Add(s1, s2),
				s3),
			fq6.F.Add(s0, s4)),
	}
}
