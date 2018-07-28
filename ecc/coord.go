package ecc

import (
	"bytes"
	"math/big"
)

var (
	bigZero   = big.NewInt(int64(0))
	zeroPoint = Point{bigZero, bigZero}
)

type Point struct {
	X *big.Int
	Y *big.Int
}

func (c1 *Point) Equal(c2 Point) bool {
	if !bytes.Equal(c1.X.Bytes(), c2.X.Bytes()) {
		return false
	}
	if !bytes.Equal(c1.Y.Bytes(), c2.Y.Bytes()) {
		return false
	}
	return true
}
