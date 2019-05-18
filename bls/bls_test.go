package bls

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBls(t *testing.T) {
	bls, err := NewBLS()
	assert.Nil(t, err)
	keys0, err := bls.NewKeys()
	assert.Nil(t, err)

	fmt.Println("privK:", keys0.PrivK)
	fmt.Println("pubK:", keys0.PubK)

	m0 := []byte("message0")
	sig0 := bls.Sign(keys0.PrivK, m0)
	fmt.Println("signature:", sig0)

	verified := bls.Verify(m0, sig0, keys0.PubK)
	fmt.Println("one signature verified:", verified)
	assert.True(t, verified)

	// signature aggregation
	keys1, err := bls.NewKeys()
	assert.Nil(t, err)
	sig1 := bls.Sign(keys1.PrivK, m0)
	assert.True(t, bls.Verify(m0, sig1, keys1.PubK))

	keys2, err := bls.NewKeys()
	assert.Nil(t, err)
	sig2 := bls.Sign(keys2.PrivK, m0)

	aggr := bls.AggregateSignatures(sig0, sig1, sig2)

	pubKArray := [][3]*big.Int{keys0.PubK, keys1.PubK, keys2.PubK}
	verified = bls.VerifyAggregatedSignatures(aggr, pubKArray, m0)
	fmt.Println("signature aggregation verified:", verified)
	assert.True(t, verified)

}
