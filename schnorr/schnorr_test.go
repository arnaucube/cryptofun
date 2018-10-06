package schnorr

import (
	"math/big"
	"testing"

	"github.com/arnaucube/cryptofun/ecc"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	c := ecc.Point{big.NewInt(int64(7)), big.NewInt(int64(8))} // Generator
	h := Hash([]byte("hola"), c)
	assert.Equal(t, h.String(), "34719153732582497359642109898768696927847420320548121616059449972754491425079")
}

func TestSign(t *testing.T) {
	ec := ecc.NewEC(big.NewInt(int64(0)), big.NewInt(int64(7)), big.NewInt(int64(11)))
	g := ecc.Point{big.NewInt(int64(7)), big.NewInt(int64(8))} // Generator
	r := big.NewInt(int64(7))                                  // random r
	schnorr, sk, err := Gen(ec, g, r)
	assert.Nil(t, err)

	m := []byte("hola")

	s, rPoint, err := schnorr.Sign(sk, m)
	assert.Nil(t, err)

	verified, err := Verify(schnorr.EC, sk.PubK, m, s, rPoint)
	assert.Nil(t, err)

	assert.True(t, verified)
}

func TestSign2(t *testing.T) {
	ec := ecc.NewEC(big.NewInt(int64(0)), big.NewInt(int64(7)), big.NewInt(int64(29)))
	g := ecc.Point{big.NewInt(int64(11)), big.NewInt(int64(27))} // Generator
	r := big.NewInt(int64(23))                                   // random r
	schnorr, sk, err := Gen(ec, g, r)
	assert.Nil(t, err)

	m := []byte("hola")

	s, rPoint, err := schnorr.Sign(sk, m)
	assert.Nil(t, err)

	verified, err := Verify(schnorr.EC, sk.PubK, m, s, rPoint)
	assert.Nil(t, err)

	assert.True(t, verified)
}
