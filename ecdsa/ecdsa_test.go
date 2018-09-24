package ecdsa

import (
	"math/big"
	"testing"

	"github.com/arnaucube/cryptofun/ecc"
	"github.com/stretchr/testify/assert"
)

func TestNewECDSA(t *testing.T) {
	ec := ecc.NewEC(1, 18, 19)
	g := ecc.Point{big.NewInt(int64(7)), big.NewInt(int64(11))}
	dsa, err := NewDSA(ec, g)
	assert.Nil(t, err)

	privK := big.NewInt(int64(5))
	pubK, err := dsa.PubK(privK)
	assert.Nil(t, err)

	if !pubK.Equal(ecc.Point{big.NewInt(int64(13)), big.NewInt(int64(9))}) {
		t.Errorf("pubK!=(13, 9)")
	}
}

func TestECDSASignAndVerify(t *testing.T) {
	ec := ecc.NewEC(1, 18, 19)
	g := ecc.Point{big.NewInt(int64(7)), big.NewInt(int64(11))}
	dsa, err := NewDSA(ec, g)
	assert.Nil(t, err)

	privK := big.NewInt(int64(5))
	pubK, err := dsa.PubK(privK)
	assert.Nil(t, err)

	hashval := big.NewInt(int64(40))
	r := big.NewInt(int64(11))

	sig, err := dsa.Sign(hashval, privK, r)
	assert.Nil(t, err)

	verified, err := dsa.Verify(hashval, sig, pubK)
	assert.True(t, verified)
}
