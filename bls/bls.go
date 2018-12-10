package bls

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
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

func (bls BLS) AggregateSignatures(signatures ...[3][2]*big.Int) [3][2]*big.Int {
	aggr := signatures[0]
	for _, sig := range signatures {
		aggr = bls.Bn.G2.Add(aggr, sig)
	}
	return aggr
}
func (bls BLS) VerifyAggregatedSignatures(aggrsig [3][2]*big.Int, pubKArray [][3]*big.Int, mArray [][]byte) bool {
	if len(pubKArray) != len(mArray) {
		fmt.Println("pubK array and msg array not with the same number of elements")
		return false
	}
	pairingGS, err := bls.Bn.Pairing(bls.Bn.G1.G, aggrsig)
	if err != nil {
		return false
	}
	pairingsMul, err := bls.Bn.Pairing(pubKArray[0], bls.Hash(mArray[0]))
	if err != nil {
		return false
	}

	for i := 1; i < len(pubKArray); i++ {
		e, err := bls.Bn.Pairing(pubKArray[i], bls.Hash(mArray[i]))
		if err != nil {
			return false
		}
		pairingsMul = bls.Bn.Fq12.Mul(pairingsMul, e)
	}
	if !bls.Bn.Fq12.Equal(pairingGS, pairingsMul) {
		return false
	}
	return true
}
