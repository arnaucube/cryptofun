package paillier

import (
	"bytes"
	"fmt"
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
	d := Decrypt(c, key.PubK, key.PrivK)
	if m == d {
		fmt.Println(key)
		t.Errorf("m not equal to decrypted")
	}
}

func TestHomomorphicAddition(t *testing.T) {
	// key, err := GenerateKeyPair()
	// assert.Nil(t, err)

	// key harcoded for tests
	pubK := PublicKey{
		N: big.NewInt(204223),
		G: big.NewInt(24929195694),
	}
	privK := PrivateKey{
		Lambda: big.NewInt(101660),
		Mu:     big.NewInt(117648),
	}
	key := Key{
		PubK:  pubK,
		PrivK: privK,
	}

	n1 := big.NewInt(int64(110))
	n2 := big.NewInt(int64(150))
	c1 := Encrypt(n1, key.PubK)
	c2 := Encrypt(n2, key.PubK)
	c3c4 := HomomorphicAddition(c1, c2, key.PubK)
	d := Decrypt(c3c4, key.PubK, key.PrivK)
	if !bytes.Equal(new(big.Int).Add(n1, n2).Bytes(), d.Bytes()) {
		fmt.Println(key)
		t.Errorf("decrypted result not equal to original result")
	}
}
