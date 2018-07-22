package secrets

import (
	"crypto/rand"
	"math/big"
	"testing"
)

func TestCreate(t *testing.T) {
	k := 123456789
	p, err := rand.Prime(rand.Reader, bits/2)
	if err != nil {
		t.Errorf(err.Error())
	}

	nNeededSecrets := big.NewInt(int64(3))
	nShares := big.NewInt(int64(6))
	shares, err := Create(
		nNeededSecrets,
		nShares,
		p,
		big.NewInt(int64(k)))
	if err != nil {
		t.Errorf(err.Error())
	}

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
	if int64(k) != secr.Int64() {
		t.Errorf("reconstructed secret not correspond to original secret")
	}
}
