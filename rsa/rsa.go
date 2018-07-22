package rsa

import (
	"bytes"
	"crypto/rand"
	"math/big"
)

const (
	bits = 512 // 2048
)

var bigOne = big.NewInt(int64(1))

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

func GenerateKeyPair() (key Key, err error) {
	p, err := rand.Prime(rand.Reader, bits/2)
	if err != nil {
		return key, err
	}
	q, err := rand.Prime(rand.Reader, bits/2)
	if err != nil {
		return key, err
	}

	n := new(big.Int).Mul(p, q)
	p_1 := new(big.Int).Sub(p, bigOne)
	q_1 := new(big.Int).Sub(q, bigOne)
	phi := new(big.Int).Mul(p_1, q_1)
	e := 65537
	var pubK PublicKey
	pubK.E = big.NewInt(int64(e))
	pubK.N = n

	d := new(big.Int).ModInverse(big.NewInt(int64(e)), phi)

	var privK PrivateKey
	privK.D = d
	privK.N = n

	key.PubK = pubK
	key.PrivK = privK
	return key, nil
}

func Encrypt(m *big.Int, pubK PublicKey) *big.Int {
	c := new(big.Int).Exp(m, pubK.E, pubK.N)
	return c
}
func Decrypt(c *big.Int, privK PrivateKey) *big.Int {
	m := new(big.Int).Exp(c, privK.D, privK.N)
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
