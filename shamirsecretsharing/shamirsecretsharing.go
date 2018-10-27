package shamirsecretsharing

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const (
	// bits = 1024
	bits = 2048
)

// Create calculates the secrets to share from given parameters
// t: number of secrets needed
// n: number of shares
// p: random point
// k: secret to share
func Create(t, n, p, k *big.Int) (result [][]*big.Int, err error) {
	if k.Cmp(p) > 0 {
		return nil, errors.New("Error: need k<p. k: " + k.String() + ", p: " + p.String())
	}
	//generate the basePolynomial
	var basePolynomial []*big.Int
	basePolynomial = append(basePolynomial, k)
	for i := 0; i < int(t.Int64())-1; i++ {
		randPrime, err := rand.Prime(rand.Reader, bits/2)
		if err != nil {
			return result, err
		}
		basePolynomial = append(basePolynomial, randPrime)
	}

	//calculate shares, based on the basePolynomial
	var shares []*big.Int
	for i := 1; i < int(n.Int64())+1; i++ {
		var pResultMod *big.Int
		pResult := big.NewInt(int64(0))
		for x, polElem := range basePolynomial {
			if x == 0 {
				pResult = pResult.Add(pResult, polElem)
			} else {
				iBigInt := big.NewInt(int64(i))
				xBigInt := big.NewInt(int64(x))
				iPowed := iBigInt.Exp(iBigInt, xBigInt, nil)
				currElem := iPowed.Mul(iPowed, polElem)
				pResult = pResult.Add(pResult, currElem)
				pResultMod = pResult.Mod(pResult, p)
			}
		}
		shares = append(shares, pResultMod)
	}
	//put the share together with his p value
	result = packSharesAndI(shares)
	return result, nil
}

func packSharesAndI(sharesString []*big.Int) (r [][]*big.Int) {
	for i, share := range sharesString {
		curr := []*big.Int{share, big.NewInt(int64(i + 1))}
		r = append(r, curr)
	}
	return r
}
func unpackSharesAndI(sharesPacked [][]*big.Int) ([]*big.Int, []*big.Int) {
	var shares []*big.Int
	var i []*big.Int
	for _, share := range sharesPacked {
		shares = append(shares, share[0])
		i = append(i, share[1])
	}
	return shares, i
}

// LagrangeInterpolation calculates the secret from given shares
func LagrangeInterpolation(sharesGiven [][]*big.Int, p *big.Int) *big.Int {
	resultN := big.NewInt(int64(0))
	resultD := big.NewInt(int64(0))

	//unpack shares
	sharesBigInt, sharesIBigInt := unpackSharesAndI(sharesGiven)

	for i := 0; i < len(sharesBigInt); i++ {
		lagrangeNumerator := big.NewInt(int64(1))
		lagrangeDenominator := big.NewInt(int64(1))
		for j := 0; j < len(sharesBigInt); j++ {
			if sharesIBigInt[i] != sharesIBigInt[j] {
				currLagrangeNumerator := sharesIBigInt[j]
				currLagrangeDenominator := new(big.Int).Sub(sharesIBigInt[j], sharesIBigInt[i])
				lagrangeNumerator = new(big.Int).Mul(lagrangeNumerator, currLagrangeNumerator)
				lagrangeDenominator = new(big.Int).Mul(lagrangeDenominator, currLagrangeDenominator)
			}
		}
		numerator := new(big.Int).Mul(sharesBigInt[i], lagrangeNumerator)
		quo := new(big.Int).Quo(numerator, lagrangeDenominator)
		if quo.Int64() != 0 {
			resultN = resultN.Add(resultN, quo)
		} else {
			resultNMULlagrangeDenominator := new(big.Int).Mul(resultN, lagrangeDenominator)
			resultN = new(big.Int).Add(resultNMULlagrangeDenominator, numerator)

			resultD = resultD.Add(resultD, lagrangeDenominator)
		}
	}

	var modinvMul *big.Int
	if resultD.Int64() != 0 {
		modinv := new(big.Int).ModInverse(resultD, p)
		modinvMul = new(big.Int).Mul(resultN, modinv)
	} else {
		modinvMul = resultN
	}
	r := new(big.Int).Mod(modinvMul, p)
	return r
}
