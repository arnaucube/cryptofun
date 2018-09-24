package dh

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	bits = 2048
)

func TestDiffieHellman(t *testing.T) {
	p, err := rand.Prime(rand.Reader, bits/2)
	assert.Nil(t, err)

	g, err := rand.Prime(rand.Reader, bits/2)
	assert.Nil(t, err)

	max, err := rand.Prime(rand.Reader, bits/2)
	assert.Nil(t, err)

	a, err := rand.Int(rand.Reader, max)
	assert.Nil(t, err)

	b, err := rand.Int(rand.Reader, max)
	assert.Nil(t, err)

	A := new(big.Int).Exp(g, a, p)
	B := new(big.Int).Exp(g, b, p)

	sA := new(big.Int).Exp(B, a, p)
	sB := new(big.Int).Exp(A, b, p)

	if sA.Int64() != sB.Int64() {
		t.Errorf("secret not equal")
	}
}
