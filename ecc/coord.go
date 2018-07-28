package ecc

import "math/big"

var (
	bigZero   = big.NewInt(int64(0))
	zeroPoint = Point{bigZero, bigZero}
)

type Point struct {
	X *big.Int
	Y *big.Int
}

func (c1 *Point) Equal(c2 Point) bool {
	if c1.X.Int64() != c2.X.Int64() {
		return false
	}
	if c1.Y.Int64() != c2.Y.Int64() {
		return false
	}
	return true
}
