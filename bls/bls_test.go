package bls

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBls(t *testing.T) {
	bls, err := NewKeys()
	assert.Nil(t, err)

	fmt.Println("privK:", bls.PrivK)
	fmt.Println("pubK:", bls.PubK)

	m := []byte("test")
	sig := bls.Sign(m)
	fmt.Println("signature:", sig)

	verified := bls.Verify(m, sig, bls.PubK)
	fmt.Println("verified:", verified)
	assert.True(t, verified)
}
