package ecdsa

import (
	"bytes"
	"math/big"

	// ecc "../ecc"
	"github.com/arnaucube/cryptofun/ecc"
)

// DSA is the ECDSA data structure
type DSA struct {
	EC ecc.EC
	G  ecc.Point
	N  *big.Int
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
func (dsa DSA) PubK(privK *big.Int) (ecc.Point, error) {
	// privK: rand < ec.Q
	privKCopy := new(big.Int).SetBytes(privK.Bytes())
	pubK, err := dsa.EC.Mul(dsa.G, privKCopy)
	return pubK, err
}

// Sign performs the ECDSA signature
func (dsa DSA) Sign(hashval *big.Int, privK *big.Int, r *big.Int) ([2]*big.Int, error) {
	rCopy := new(big.Int).SetBytes(r.Bytes())
	m, err := dsa.EC.Mul(dsa.G, rCopy)
	if err != nil {
		return [2]*big.Int{}, err
	}
	// inv(r) mod dsa.N
	inv := new(big.Int).ModInverse(r, dsa.N)
	// m.X * privK
	privKCopy := new(big.Int).SetBytes(privK.Bytes())
	xPrivK := new(big.Int).Mul(m.X, privKCopy)
	// (hashval + m.X * privK)
	hashvalXPrivK := new(big.Int).Add(hashval, xPrivK)
	// inv * (hashval + m.X * privK) mod dsa.N
	a := new(big.Int).Mul(inv, hashvalXPrivK)
	r2 := new(big.Int).Mod(a, dsa.N)
	return [2]*big.Int{m.X, r2}, err
}

// Verify validates the ECDSA signature
func (dsa DSA) Verify(hashval *big.Int, sig [2]*big.Int, pubK ecc.Point) (bool, error) {
	w := new(big.Int).ModInverse(sig[1], dsa.N)
	wCopy := new(big.Int).SetBytes(w.Bytes())
	u1raw := new(big.Int).Mul(hashval, wCopy)
	u1 := new(big.Int).Mod(u1raw, dsa.N)
	wCopy = new(big.Int).SetBytes(w.Bytes())
	u2raw := new(big.Int).Mul(sig[0], wCopy)
	u2 := new(big.Int).Mod(u2raw, dsa.N)

	gU1, err := dsa.EC.Mul(dsa.G, u1)
	if err != nil {
		return false, err
	}
	pubKU2, err := dsa.EC.Mul(pubK, u2)
	if err != nil {
		return false, err
	}
	p, err := dsa.EC.Add(gU1, pubKU2)
	if err != nil {
		return false, err
	}
	pXmodN := new(big.Int).Mod(p.X, dsa.N)
	return bytes.Equal(pXmodN.Bytes(), sig[0].Bytes()), nil
}
