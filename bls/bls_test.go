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
	m1 := []byte("message1")
	sig1 := bls.Sign(keys1.PrivK, m1)

	keys2, err := bls.NewKeys()
	assert.Nil(t, err)
	m2 := []byte("message2")
	sig2 := bls.Sign(keys2.PrivK, m2)

	aggr := bls.AggregateSignatures(sig0, sig1, sig2)

	pubKArray := [][3]*big.Int{keys0.PubK, keys1.PubK, keys2.PubK}
	mArray := [][]byte{m0, m1, m2}
	verified = bls.VerifyAggregatedSignatures(aggr, pubKArray, mArray)
	fmt.Println("signature aggregation verified:", verified)
	assert.True(t, verified)

}
