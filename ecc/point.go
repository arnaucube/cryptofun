package ecc

import (
	"bytes"
	"math/big"
)

var (
	bigZero   = big.NewInt(int64(0))
	bigOne    = big.NewInt(int64(1))
	zeroPoint = Point{bigZero, bigZero}
)

// Point is the data structure for a point, containing the X and Y coordinates
type Point struct {
	X *big.Int
	Y *big.Int
}

// Equal compares the X and Y coord of a Point and returns true if are the same
func (c1 *Point) Equal(c2 Point) bool {
	if !bytes.Equal(c1.X.Bytes(), c2.X.Bytes()) {
		return false
	}
	if !bytes.Equal(c1.Y.Bytes(), c2.Y.Bytes()) {
		return false
	}
	return true
}

// String returns the components of the point in a string
func (p *Point) String() string {
	return "(" + p.X.String() + ", " + p.Y.String() + ")"
}
