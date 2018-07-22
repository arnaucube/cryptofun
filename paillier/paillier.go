package paillier

import (
	"crypto/rand"
	"errors"
	"math/big"

	prime "../prime"
)

const (
	bits = 16
)

type PublicKey struct {
	N *big.Int `json:"n"`
	G *big.Int `json:"g"`
}
type PrivateKey struct {
	Lambda *big.Int `json:"lambda"`
	Mu     *big.Int `json:"mu"`
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

	pq := new(big.Int).Mul(p, q)
	p1q1 := big.NewInt((p.Int64() - 1) * (q.Int64() - 1))
	gcd := new(big.Int).GCD(nil, nil, pq, p1q1)
	if gcd.Int64() != int64(1) {
		return key, errors.New("gcd comprovation failed")
	}

	n := new(big.Int).Mul(p, q)
	lambda := big.NewInt(int64(Lcm(float64(p.Int64())-1, float64(q.Int64())-1)))

	//g generation
	alpha := big.NewInt(int64(prime.RandInt(0, int(n.Int64()))))
	beta := big.NewInt(int64(prime.RandInt(0, int(n.Int64()))))
	alphan := new(big.Int).Mul(alpha, n)
	alphan1 := new(big.Int).Add(alphan, big.NewInt(1))
	betaN := new(big.Int).Exp(beta, n, nil)
	ab := new(big.Int).Mul(alphan1, betaN)
	n2 := new(big.Int).Mul(n, n)
	g := new(big.Int).Mod(ab, n2)
	//in some Paillier implementations use this:
	// g = new(big.Int).Add(n, big.NewInt(1))

	key.PubK.N = n
	key.PubK.G = g

	//mu generation
	Glambda := new(big.Int).Exp(g, lambda, nil)
	u := new(big.Int).Mod(Glambda, n2)
	L := L(u, n)
	mu := new(big.Int).ModInverse(L, n)

	key.PrivK.Lambda = lambda
	key.PrivK.Mu = mu

	return key, nil
}

func Lcm(a, b float64) float64 {
	r := (a * b) / float64(prime.Gcd(int(a), int(b)))
	return r

}
func L(u *big.Int, n *big.Int) *big.Int {
	u1 := new(big.Int).Sub(u, big.NewInt(1))
	L := new(big.Int).Div(u1, n)
	return L
}

func Encrypt(m *big.Int, pubK PublicKey) *big.Int {
	gM := new(big.Int).Exp(pubK.G, m, nil)
	r := big.NewInt(int64(prime.RandInt(0, int(pubK.N.Int64()))))
	rN := new(big.Int).Exp(r, pubK.N, nil)
	gMrN := new(big.Int).Mul(gM, rN)
	n2 := new(big.Int).Mul(pubK.N, pubK.N)
	c := new(big.Int).Mod(gMrN, n2)
	return c
}
func Decrypt(c *big.Int, pubK PublicKey, privK PrivateKey) *big.Int {
	cLambda := new(big.Int).Exp(c, privK.Lambda, nil)
	n2 := new(big.Int).Mul(pubK.N, pubK.N)
	u := new(big.Int).Mod(cLambda, n2)
	L := L(u, pubK.N)
	LMu := new(big.Int).Mul(L, privK.Mu)
	m := new(big.Int).Mod(LMu, pubK.N)
	return m
}

func HomomorphicAddition(c1 *big.Int, c2 *big.Int, pubK PublicKey) *big.Int {
	c1c2 := new(big.Int).Mul(c1, c2)
	n2 := new(big.Int).Mul(pubK.N, pubK.N)
	d := new(big.Int).Mod(c1c2, n2)
	return d
}
