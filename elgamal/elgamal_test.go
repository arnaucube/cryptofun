package elgamal

import (
	"math/big"
	"testing"

	"github.com/arnaucube/cryptofun/ecc"
	"github.com/stretchr/testify/assert"
)

func TestNewEG(t *testing.T) {
	ec := ecc.NewEC(big.NewInt(int64(1)), big.NewInt(int64(18)), big.NewInt(int64(19)))
	g := ecc.Point{big.NewInt(int64(7)), big.NewInt(int64(11))}
	eg, err := NewEG(ec, g)
	assert.Nil(t, err)

	privK := big.NewInt(int64(5))
	pubK, err := eg.PubK(privK)
	assert.Nil(t, err)

	if !pubK.Equal(ecc.Point{big.NewInt(int64(13)), big.NewInt(int64(9))}) {
		t.Errorf("pubK!=(13, 9)")
	}
}
func TestEGEncrypt(t *testing.T) {
	ec := ecc.NewEC(big.NewInt(int64(1)), big.NewInt(int64(18)), big.NewInt(int64(19)))
	g := ecc.Point{big.NewInt(int64(7)), big.NewInt(int64(11))}
	eg, err := NewEG(ec, g)
	assert.Nil(t, err)

	privK := big.NewInt(int64(5))
	pubK, err := eg.PubK(privK)
	assert.Nil(t, err)

	// m: point to encrypt
	m := ecc.Point{big.NewInt(int64(11)), big.NewInt(int64(12))}
	c, err := eg.Encrypt(m, pubK, big.NewInt(int64(15)))
	assert.Nil(t, err)

	if !c[0].Equal(ecc.Point{big.NewInt(int64(8)), big.NewInt(int64(5))}) {
		t.Errorf("c[0] != (8, 5), encryption failed")
	}
	if !c[1].Equal(ecc.Point{big.NewInt(int64(2)), big.NewInt(int64(16))}) {
		t.Errorf("c[1] != (2, 16), encryption failed")
	}
}

func TestEGDecrypt(t *testing.T) {
	ec := ecc.NewEC(big.NewInt(int64(1)), big.NewInt(int64(18)), big.NewInt(int64(19)))
	g := ecc.Point{big.NewInt(int64(7)), big.NewInt(int64(11))}
	eg, err := NewEG(ec, g)
	assert.Nil(t, err)

	privK := big.NewInt(int64(5))
	pubK, err := eg.PubK(privK)
	assert.Nil(t, err)

	// m: point to encrypt
	m := ecc.Point{big.NewInt(int64(11)), big.NewInt(int64(12))}
	c, err := eg.Encrypt(m, pubK, big.NewInt(int64(15)))
	assert.Nil(t, err)

	d, err := eg.Decrypt(c, privK)
	assert.Nil(t, err)

	if !m.Equal(d) {
		t.Errorf("m != d, decrypting failed")
	}
}
