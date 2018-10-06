package ecc

import (
	"bytes"
	"math/big"
)

var (
	BigZero   = big.NewInt(int64(0))
	BigOne    = big.NewInt(int64(1))
	ZeroPoint = Point{BigZero, BigZero}
)

// Point is the data structure for a point, containing the X and Y coordinates
type Point struct {
	X *big.Int
	Y *big.Int
}

// Equal compares the X and Y coord of a Point and returns true if are the same
func (p1 *Point) Equal(p2 Point) bool {
	if !bytes.Equal(p1.X.Bytes(), p2.X.Bytes()) {
		return false
	}
	if !bytes.Equal(p1.Y.Bytes(), p2.Y.Bytes()) {
		return false
	}
	return true
}

// String returns the components of the point in a string
func (p *Point) String() string {
	return "(" + p.X.String() + ", " + p.Y.String() + ")"
}
