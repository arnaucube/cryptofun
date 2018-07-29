package ecdsa

import (
	"bytes"
	"math/big"

	ecc "../ecc"
)

// DSA is the ECDSA data structure
type DSA struct {
	EC ecc.EC
	G  ecc.Point
	N  int
}

// NewDSA defines a new DSA data structure
func NewDSA(ec ecc.EC, g ecc.Point) (DSA, error) {
	var dsa DSA
	var err error
	dsa.EC = ec
	dsa.G = g
	dsa.N, err = ec.Order(g)
	return dsa, err
}

// PubK returns the public key Point calculated from the private key over the elliptic curve
func (dsa DSA) PubK(privK int) (ecc.Point, error) {
	// privK: rand < ec.Q
	pubK, err := dsa.EC.Mul(dsa.G, privK)
	return pubK, err
}
func (dsa DSA) Sign(hashval *big.Int, privK int, r *big.Int) ([2]*big.Int, error) {
	m, err := dsa.EC.Mul(dsa.G, int(r.Int64()))
	if err != nil {
		return [2]*big.Int{}, err
	}
	// inv(r) mod dsa.N
	inv := new(big.Int).ModInverse(r, big.NewInt(int64(dsa.N)))
	// m.X * privK
	xPrivK := new(big.Int).Mul(m.X, big.NewInt(int64(privK)))
	// (hashval + m.X * privK)
	hashvalXPrivK := new(big.Int).Add(hashval, xPrivK)
	// inv * (hashval + m.X * privK) mod dsa.N
	a := new(big.Int).Mul(inv, hashvalXPrivK)
	r2 := new(big.Int).Mod(a, big.NewInt(int64(dsa.N)))
	return [2]*big.Int{m.X, r2}, err
}

func (dsa DSA) Verify(hashval *big.Int, sig [2]*big.Int, pubK ecc.Point) (bool, error) {
	w := new(big.Int).ModInverse(sig[1], big.NewInt(int64(dsa.N)))
	u1raw := new(big.Int).Mul(hashval, w)
	u1 := new(big.Int).Mod(u1raw, big.NewInt(int64(dsa.N)))
	u2raw := new(big.Int).Mul(sig[0], w)
	u2 := new(big.Int).Mod(u2raw, big.NewInt(int64(dsa.N)))

	gU1, err := dsa.EC.Mul(dsa.G, int(u1.Int64()))
	if err != nil {
		return false, err
	}
	pubKU2, err := dsa.EC.Mul(pubK, int(u2.Int64()))
	if err != nil {
		return false, err
	}
	p, err := dsa.EC.Add(gU1, pubKU2)
	if err != nil {
		return false, err
	}
	pXmodN := new(big.Int).Mod(p.X, big.NewInt(int64(dsa.N)))
	return bytes.Equal(pXmodN.Bytes(), sig[0].Bytes()), nil
}
