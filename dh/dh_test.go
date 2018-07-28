package dh

import (
	"crypto/rand"
	"math/big"
	"testing"
)

const (
	bits = 2048
)

func TestDiffieHellman(t *testing.T) {
	p, err := rand.Prime(rand.Reader, bits/2)
	if err != nil {
		t.Errorf(err.Error())
	}
	g, err := rand.Prime(rand.Reader, bits/2)
	if err != nil {
		t.Errorf(err.Error())
	}

	max, err := rand.Prime(rand.Reader, bits/2)
	if err != nil {
		t.Errorf(err.Error())
	}
	a, err := rand.Int(rand.Reader, max)
	if err != nil {
		t.Errorf(err.Error())
	}
	b, err := rand.Int(rand.Reader, max)
	if err != nil {
		t.Errorf(err.Error())
	}

	A := new(big.Int).Exp(g, a, p)
	B := new(big.Int).Exp(g, b, p)

	sA := new(big.Int).Exp(B, a, p)
	sB := new(big.Int).Exp(A, b, p)

	if sA.Int64() != sB.Int64() {
		t.Errorf("secret not equal")
	}
}
