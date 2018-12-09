package bls

import (
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"github.com/arnaucube/go-snark/bn128"
)

// this BLS implementation uses the Go implementation of the BN128 pairing github.com/arnaucube/go-snark/bn128

const bits = 2048

// BLS is the data structure of the BLS signature scheme, including the BN128 pairing curve
type BLS struct {
	Bn    bn128.Bn128
	PrivK *big.Int
	PubK  [3]*big.Int
}

// NewKeys generate new Private Key and Public Key
func NewKeys() (BLS, error) {
	bn, err := bn128.NewBn128()
	if err != nil {
		return BLS{}, err
	}
	bls := BLS{}
	bls.Bn = bn
	bls.PrivK, err = rand.Prime(rand.Reader, bits)
	if err != nil {
		return BLS{}, err
	}

	// pubK = pk * G
	bls.PubK = bls.Bn.G1.MulScalar(bls.Bn.G1.G, bls.PrivK)
	return bls, nil
}

// Hash hashes a message m
func (bls BLS) Hash(m []byte) [3][2]*big.Int {
	h := sha256.New()
	h.Write(m)
	hash := h.Sum(nil)
	r := new(big.Int).SetBytes(hash)
	// get point over the curve
	point := bls.Bn.G2.MulScalar(bls.Bn.G2.G, r)
	return point
}

// Sign performs the BLS signature of a message m
func (bls BLS) Sign(m []byte) [3][2]*big.Int {
	// s = pk * H(m)
	h := bls.Hash(m)
	sig := bls.Bn.G2.MulScalar(h, bls.PrivK)
	return sig
}

// Verify checks the signature of a message m with the given Public Key
func (bls BLS) Verify(m []byte, sig [3][2]*big.Int, pubK [3]*big.Int) bool {
	// checks e(P, G) == e(G, s)
	p1, err := bls.Bn.Pairing(bls.PubK, bls.Hash(m))
	if err != nil {
		return false
	}
	p2, err := bls.Bn.Pairing(bls.Bn.G1.G, sig)
	if err != nil {
		return false
	}

	return bls.Bn.Fq12.Equal(p1, p2)
}
