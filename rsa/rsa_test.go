package rsa

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptDecrypt(t *testing.T) {
	key, err := GenerateKeyPair()
	assert.Nil(t, err)

	mBytes := []byte("Hi")
	m := new(big.Int).SetBytes(mBytes)
	c := Encrypt(m, key.PubK)

	d := Decrypt(c, key.PrivK)
	if m == d {
		t.Errorf("m not equal to decrypted")
	}
}
func TestBlindSignature(t *testing.T) {
	key, err := GenerateKeyPair()
	assert.Nil(t, err)

	mBytes := []byte("Hi")
	m := new(big.Int).SetBytes(mBytes)
	c := Encrypt(m, key.PubK)

	d := Decrypt(c, key.PrivK)
	if m == d {
		t.Errorf("decrypted d not equal to original m")
	}
	rVal := big.NewInt(int64(101))
	mBlinded := Blind(m, rVal, key.PubK)
	sigma := BlindSign(mBlinded, key.PrivK)
	mSigned := Unblind(sigma, rVal, key.PubK)
	verified := Verify(m, mSigned, key.PubK)
	if !verified {
		t.Errorf("false, signature not verified")
	}
}

func TestHomomorphicMultiplication(t *testing.T) {
	key, err := GenerateKeyPair()
	assert.Nil(t, err)

	n1 := big.NewInt(int64(11))
	n2 := big.NewInt(int64(15))
	c1 := Encrypt(n1, key.PubK)
	c2 := Encrypt(n2, key.PubK)
	c3c4 := HomomorphicMul(c1, c2, key.PubK)
	d := Decrypt(c3c4, key.PrivK)
	if !bytes.Equal(new(big.Int).Mul(n1, n2).Bytes(), d.Bytes()) {
		t.Errorf("decrypted result not equal to original result")
	}
}
