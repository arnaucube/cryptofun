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
	Bn bn128.Bn128
}
type BLSKeys struct {
	PrivK *big.Int
	PubK  [3]*big.Int
}

// NewBLS generates a new BLS scheme
func NewBLS() (BLS, error) {
	bn, err := bn128.NewBn128()
	if err != nil {
		return BLS{}, err
	}
	bls := BLS{}
	bls.Bn = bn
	return bls, nil
}

// NewKeys generate new Private Key and Public Key
func (bls BLS) NewKeys() (BLSKeys, error) {
	var err error
	k := BLSKeys{}
	k.PrivK, err = rand.Prime(rand.Reader, bits)
	if err != nil {
		return BLSKeys{}, err
	}

	// pubK = pk * G
	k.PubK = bls.Bn.G1.MulScalar(bls.Bn.G1.G, k.PrivK)
	return k, nil
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
func (bls BLS) Sign(privK *big.Int, m []byte) [3][2]*big.Int {
	// s = pk * H(m)
	h := bls.Hash(m)
	sig := bls.Bn.G2.MulScalar(h, privK)
	return sig
}

// Verify checks the signature of a message m with the given Public Key
func (bls BLS) Verify(m []byte, sig [3][2]*big.Int, pubK [3]*big.Int) bool {
	// checks e(P, G) == e(G, s)
	p1, err := bls.Bn.Pairing(pubK, bls.Hash(m))
	if err != nil {
		return false
	}
	p2, err := bls.Bn.Pairing(bls.Bn.G1.G, sig)
	if err != nil {
		return false
	}

	return bls.Bn.Fq12.Equal(p1, p2)
}

// AggregateSignatures
// s = s0 + s1 + s2 ...
func (bls BLS) AggregateSignatures(signatures ...[3][2]*big.Int) [3][2]*big.Int {
	aggr := signatures[0]
	for i := 1; i < len(signatures); i++ {
		aggr = bls.Bn.G2.Add(aggr, signatures[i])
	}
	return aggr
}

// VerifyAggregatedSignatures
// ê(G,S) == ê(P, H(m))
// ê(G, s0+s1+s2...) == ê(p0+p1+p2..., H(m))
func (bls BLS) VerifyAggregatedSignatures(aggrsig [3][2]*big.Int, pubKArray [][3]*big.Int, m []byte) bool {
	aggrPubKs := pubKArray[0]
	for i := 1; i < len(pubKArray); i++ {
		aggrPubKs = bls.Bn.G1.Add(aggrPubKs, pubKArray[i])
	}

	left, err := bls.Bn.Pairing(bls.Bn.G1.G, aggrsig)
	if err != nil {
		return false
	}

	right, err := bls.Bn.Pairing(aggrPubKs, bls.Hash(m))
	if err != nil {
		return false
	}

	if !bls.Bn.Fq12.Equal(left, right) {
		return false
	}
	return true
}
