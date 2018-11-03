package bn128

import (
	"bytes"
	"math/big"
)

// Fq is the Z field over modulus Q
type Fq struct {
	Q *big.Int // Q
}

// NewFq generates a new Fq
func NewFq(q *big.Int) Fq {
	return Fq{
		q,
	}
}

// Zero returns a Zero value on the Fq
func (fq Fq) Zero() *big.Int {
	return big.NewInt(int64(0))
}

// One returns a One value on the Fq
func (fq Fq) One() *big.Int {
	return big.NewInt(int64(1))
}

// Add performs an addition on the Fq
func (fq Fq) Add(a, b *big.Int) *big.Int {
	sum := new(big.Int).Add(a, b)
	return new(big.Int).Mod(sum, fq.Q)
}

// Double performs a doubling on the Fq
func (fq Fq) Double(a *big.Int) *big.Int {
	sum := new(big.Int).Add(a, a)
	return new(big.Int).Mod(sum, fq.Q)
}

// Sub performs a substraction on the Fq
func (fq Fq) Sub(a, b *big.Int) *big.Int {
	sum := new(big.Int).Sub(a, b)
	return new(big.Int).Mod(sum, fq.Q)
}

// Neg performs a negation on the Fq
func (fq Fq) Neg(a *big.Int) *big.Int {
	m := new(big.Int).Neg(a)
	return new(big.Int).Mod(m, fq.Q)
}

// Mul performs a multiplication on the Fq
func (fq Fq) Mul(a, b *big.Int) *big.Int {
	m := new(big.Int).Mul(a, b)
	return new(big.Int).Mod(m, fq.Q)
}

func (fq Fq) MulScalar(base, e *big.Int) *big.Int {
	return fq.Mul(base, e)
}

// Inverse returns the inverse on the Fq
func (fq Fq) Inverse(a *big.Int) *big.Int {
	return new(big.Int).ModInverse(a, fq.Q)
}

// Div performs a division on the Fq
func (fq Fq) Div(a, b *big.Int) *big.Int {
	// not used in fq1, method added to fit the interface
	return a
}

// Square performs a square operation on the Fq
func (fq Fq) Square(a *big.Int) *big.Int {
	m := new(big.Int).Mul(a, a)
	return new(big.Int).Mod(m, fq.Q)
}

func (fq Fq) IsZero(a *big.Int) bool {
	return bytes.Equal(a.Bytes(), fq.Zero().Bytes())
}

func (fq Fq) Copy(a *big.Int) *big.Int {
	return new(big.Int).SetBytes(a.Bytes())
}

func (fq Fq) Affine(a *big.Int) *big.Int {
	return a
}
func (fq Fq) Equal(a, b *big.Int) bool {
	return bytes.Equal(a.Bytes(), b.Bytes())
}
