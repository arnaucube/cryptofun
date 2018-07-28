package ecc

import (
	"errors"
	"math/big"
)

type EC struct {
	A *big.Int
	B *big.Int
	Q *big.Int
}

/*
	(y^2 = x^3 + Ax + B ) mod Q
	Q: prime number
*/
func NewEC(a, b, q int) (ec EC) {
	ec.A = big.NewInt(int64(a))
	ec.B = big.NewInt(int64(b))
	ec.Q = big.NewInt(int64(q))
	return ec
}

// At gets a point x in the curve
func (ec *EC) At(x *big.Int) (Point, Point, error) {
	if x.Cmp(ec.Q) > 0 {
		return Point{}, Point{}, errors.New("x<ec.Q")
	}
	// y^2 = (x^3 + ax + b) mod q
	// y = sqrt (x^3 + ax + b) mod q
	// x^3
	x3 := new(big.Int).Exp(x, big.NewInt(int64(3)), nil)
	// a^x
	aX := new(big.Int).Mul(ec.A, x)
	// x^3 + a^x
	x3aX := new(big.Int).Add(x3, aX)
	// x^3 + a^x + b
	x3aXb := new(big.Int).Add(x3aX, ec.B)
	// y = sqrt (x^3 + ax + b) mod q
	y := new(big.Int).ModSqrt(x3aXb, ec.Q)
	return Point{x, y}, Point{x, new(big.Int).Sub(ec.Q, y)}, nil
}

// TODO add valid checker point function

func (ec *EC) Neg(p Point) Point {
	// TODO get error when point not found on the ec
	return Point{p.X, new(big.Int).Sub(ec.Q, p.Y)}
}

// Add adds two points p1 and p2 and gets q
func (ec *EC) Add(p1, p2 Point) (Point, error) {
	if p1.Equal(zeroPoint) {
		return p2, errors.New("p1==(0, 0)")
	}
	if p2.Equal(zeroPoint) {
		return p1, errors.New("p1==(0, 0)")
	}
	// slope
	numerator := new(big.Int).Sub(p1.Y, p2.Y)
	denominator := new(big.Int).Sub(p1.X, p2.X)
	s := new(big.Int).Div(numerator, denominator)
	// q: new point
	var q Point
	// s^2
	s2 := new(big.Int).Exp(s, big.NewInt(int64(2)), nil)
	// s^2 - p1.X
	x2Xo := new(big.Int).Sub(s2, p1.X)
	// s^2 - p1.X - p2.X
	x2XoX2 := new(big.Int).Sub(x2Xo, p2.X)
	q.X = new(big.Int).Mod(x2XoX2, ec.Q)

	// p1.X - q.X
	xoX2 := new(big.Int).Sub(p1.X, q.X)
	// s(p1.X - q.X)
	sXoX2 := new(big.Int).Mul(s, xoX2)
	// s(p1.X - q.X) - p1.Y
	sXoX2Y := new(big.Int).Sub(sXoX2, p1.Y)
	q.Y = new(big.Int).Mod(sXoX2Y, ec.Q)

	return q, nil
}
