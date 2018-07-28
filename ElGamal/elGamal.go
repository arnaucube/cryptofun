package elgamal

import (
	ecc "../ecc"
)

// EG is the ElGamal data structure
type EG struct {
	EC ecc.EC
	G  ecc.Point
	N  int
}

// NewEG defines a new EG data structure
func NewEG(ec ecc.EC, g ecc.Point) (EG, error) {
	var eg EG
	var err error
	eg.EC = ec
	eg.G = g
	eg.N, err = ec.Order(g)
	return eg, err
}

// PubK returns the public key Point calculated from the private key over the elliptic curve
func (eg EG) PubK(privK int) (ecc.Point, error) {
	// privK: rand < ec.Q
	pubK, err := eg.EC.Mul(eg.G, privK)
	return pubK, err
}

// Encrypt encrypts a point m with the public key point, returns two points
func (eg EG) Encrypt(m ecc.Point, pubK ecc.Point, r int) ([2]ecc.Point, error) {
	p1, err := eg.EC.Mul(eg.G, r)
	if err != nil {
		return [2]ecc.Point{}, err
	}
	p2, err := eg.EC.Mul(pubK, r)
	if err != nil {
		return [2]ecc.Point{}, err
	}
	p3, err := eg.EC.Add(m, p2)
	if err != nil {
		return [2]ecc.Point{}, err
	}
	c := [2]ecc.Point{p1, p3}
	return c, err
}

// Decrypt decrypts c (two points) with the private key, returns the point decrypted
func (eg EG) Decrypt(c [2]ecc.Point, privK int) (ecc.Point, error) {
	c1 := c[0]
	c2 := c[1]
	c1PrivK, err := eg.EC.Mul(c1, privK)
	if err != nil {
		return ecc.Point{}, err
	}
	c1PrivKNeg := eg.EC.Neg(c1PrivK)
	d, err := eg.EC.Add(c2, c1PrivKNeg)
	return d, err
}
