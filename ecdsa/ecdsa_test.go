package ecdsa

import (
	"math/big"
	"testing"

	ecc "../ecc"
)

func TestNewECDSA(t *testing.T) {
	ec := ecc.NewEC(1, 18, 19)
	g := ecc.Point{big.NewInt(int64(7)), big.NewInt(int64(11))}
	dsa, err := NewDSA(ec, g)
	if err != nil {
		t.Errorf(err.Error())
	}
	privK := 5
	pubK, err := dsa.PubK(privK)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !pubK.Equal(ecc.Point{big.NewInt(int64(13)), big.NewInt(int64(9))}) {
		t.Errorf("pubK!=(13, 9)")
	}
}

func TestECDSASignAndVerify(t *testing.T) {
	ec := ecc.NewEC(1, 18, 19)
	g := ecc.Point{big.NewInt(int64(7)), big.NewInt(int64(11))}
	dsa, err := NewDSA(ec, g)
	if err != nil {
		t.Errorf(err.Error())
	}
	privK := 5
	pubK, err := dsa.PubK(privK)
	if err != nil {
		t.Errorf(err.Error())
	}
	hashval := big.NewInt(int64(40))
	r := big.NewInt(int64(11))

	sig, err := dsa.Sign(hashval, privK, r)
	if err != nil {
		t.Errorf(err.Error())
	}

	verified, err := dsa.Verify(hashval, sig, pubK)
	if !verified {
		t.Errorf("verified == false")
	}
}
