package elgamal

import (
	"math/big"
	"testing"

	ecc "../ecc"
)

func TestNewEG(t *testing.T) {
	ec := ecc.NewEC(1, 18, 19)
	g := ecc.Point{big.NewInt(int64(7)), big.NewInt(int64(11))}
	eg, err := NewEG(ec, g)
	if err != nil {
		t.Errorf(err.Error())
	}
	privK := 5
	pubK, err := eg.PubK(privK)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !pubK.Equal(ecc.Point{big.NewInt(int64(13)), big.NewInt(int64(9))}) {
		t.Errorf("pubK!=(13, 9)")
	}
}
func TestEGEncrypt(t *testing.T) {
	ec := ecc.NewEC(1, 18, 19)
	g := ecc.Point{big.NewInt(int64(7)), big.NewInt(int64(11))}
	eg, err := NewEG(ec, g)
	if err != nil {
		t.Errorf(err.Error())
	}
	privK := 5
	pubK, err := eg.PubK(privK)
	if err != nil {
		t.Errorf(err.Error())
	}
	// m: point to encrypt
	m := ecc.Point{big.NewInt(int64(11)), big.NewInt(int64(12))}
	c, err := eg.Encrypt(m, pubK, 15)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !c[0].Equal(ecc.Point{big.NewInt(int64(8)), big.NewInt(int64(5))}) {
		t.Errorf("c[0] != (8, 5), encryption failed")
	}
	if !c[1].Equal(ecc.Point{big.NewInt(int64(2)), big.NewInt(int64(16))}) {
		t.Errorf("c[1] != (2, 16), encryption failed")
	}
}

func TestEGDecrypt(t *testing.T) {
	ec := ecc.NewEC(1, 18, 19)
	g := ecc.Point{big.NewInt(int64(7)), big.NewInt(int64(11))}
	eg, err := NewEG(ec, g)
	if err != nil {
		t.Errorf(err.Error())
	}
	privK := 5
	pubK, err := eg.PubK(privK)
	if err != nil {
		t.Errorf(err.Error())
	}
	// m: point to encrypt
	m := ecc.Point{big.NewInt(int64(11)), big.NewInt(int64(12))}
	c, err := eg.Encrypt(m, pubK, 15)
	if err != nil {
		t.Errorf(err.Error())
	}
	d, err := eg.Decrypt(c, privK)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !m.Equal(d) {
		t.Errorf("m != d, decrypting failed")
	}
}
