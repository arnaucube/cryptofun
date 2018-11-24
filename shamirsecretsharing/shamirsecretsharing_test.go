package shamirsecretsharing

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	k := big.NewInt(int64(123456789))
	p, err := rand.Prime(rand.Reader, bits/2)
	assert.Nil(t, err)

	nShares := big.NewInt(int64(6))
	nNeededShares := big.NewInt(int64(3))
	shares, err := Create(
		nNeededShares,
		nShares,
		p,
		k)
	assert.Nil(t, err)

	//generate sharesToUse
	var sharesToUse [][]*big.Int
	sharesToUse = append(sharesToUse, shares[2])
	sharesToUse = append(sharesToUse, shares[1])
	sharesToUse = append(sharesToUse, shares[0])
	secr := LagrangeInterpolation(sharesToUse, p)

	// fmt.Print("original secret: ")
	// fmt.Println(k)
	// fmt.Print("p: ")
	// fmt.Println(p)
	// fmt.Print("shares: ")
	// fmt.Println(shares)
	// fmt.Print("secret result: ")
	// fmt.Println(secr)
	if !bytes.Equal(k.Bytes(), secr.Bytes()) {
		t.Errorf("reconstructed secret not correspond to original secret")
	}
}
