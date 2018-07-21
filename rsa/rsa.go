package rsa

import (
	"bytes"
	"math/big"
	"math/rand"
	"time"

	prime "../prime"
)

const (
	MaxPrime = 2000
	MinPrime = 500
)

type PublicKey struct {
	E *big.Int `json:"e"`
	N *big.Int `json:"n"`
}
type PublicKeyString struct {
	E string `json:"e"`
	N string `json:"n"`
}
type PrivateKey struct {
	D *big.Int `json:"d"`
	N *big.Int `json:"n"`
}

type Key struct {
	PubK  PublicKey
	PrivK PrivateKey
}

func GenerateKeyPair() (key Key) {
	rand.Seed(time.Now().Unix())
	p := prime.RandPrime(MinPrime, MaxPrime)
	q := prime.RandPrime(MinPrime, MaxPrime)

	n := p * q
	phi := (p - 1) * (q - 1)
	e := 65537
	var pubK PublicKey
	pubK.E = big.NewInt(int64(e))
	pubK.N = big.NewInt(int64(n))

	d := new(big.Int).ModInverse(big.NewInt(int64(e)), big.NewInt(int64(phi)))

	var privK PrivateKey
	privK.D = d
	privK.N = big.NewInt(int64(n))

	key.PubK = pubK
	key.PrivK = privK
	return key
}

func Encrypt(m *big.Int, pubK PublicKey) *big.Int {
	Me := new(big.Int).Exp(m, pubK.E, nil)
	c := new(big.Int).Mod(Me, pubK.N)
	return c
}
func Decrypt(c *big.Int, privK PrivateKey) *big.Int {
	Cd := new(big.Int).Exp(c, privK.D, nil)
	m := new(big.Int).Mod(Cd, privK.N)
	return m
}

func Blind(m *big.Int, r *big.Int, pubK PublicKey) *big.Int {
	rE := new(big.Int).Exp(r, pubK.E, nil)
	mrE := new(big.Int).Mul(m, rE)
	mBlinded := new(big.Int).Mod(mrE, pubK.N)
	return mBlinded
}

func BlindSign(m *big.Int, privK PrivateKey) *big.Int {
	sigma := new(big.Int).Exp(m, privK.D, privK.N)
	return sigma
}
func Unblind(sigma *big.Int, r *big.Int, pubK PublicKey) *big.Int {
	r1 := new(big.Int).ModInverse(r, pubK.N)
	bsr := new(big.Int).Mul(sigma, r1)
	sig := new(big.Int).Mod(bsr, pubK.N)
	return sig
}
func Verify(msg *big.Int, mSigned *big.Int, pubK PublicKey) bool {
	//decrypt the mSigned with pubK
	Cd := new(big.Int).Exp(mSigned, pubK.E, nil)
	m := new(big.Int).Mod(Cd, pubK.N)
	return bytes.Equal(msg.Bytes(), m.Bytes())
}

func HomomorphicMul(c1 *big.Int, c2 *big.Int, pubK PublicKey) *big.Int {
	c1c2 := new(big.Int).Mul(c1, c2)
	n2 := new(big.Int).Mul(pubK.N, pubK.N)
	d := new(big.Int).Mod(c1c2, n2)
	return d
}
