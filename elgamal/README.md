## ECC ElGamal
- https://en.wikipedia.org/wiki/ElGamal_encryption

- [x] ECC ElGamal key generation
- [x] ECC ElGamal Encrypton
- [x] ECC ElGamal Decryption


#### Usage
- NewEG, Encryption, Decryption
```go
// define new elliptic curve
ec := ecc.NewEC(big.NewInt(int64(1)), big.NewInt(int64(18)), big.NewInt(int64(19)))

// define new point
g := ecc.Point{big.NewInt(int64(7)), big.NewInt(int64(11))}

// define new ElGamal crypto system with the elliptic curve and the point
eg, err := NewEG(ec, g)
if err!=nil {
	fmt.Println(err)
}

// define privK&pubK over the elliptic curve
privK := big.NewInt(int64(5))
pubK, err := eg.PubK(privK)
if err!=nil {
	fmt.Println(err)
}

// define point to encrypt
m := ecc.Point{big.NewInt(int64(11)), big.NewInt(int64(12))}

// encrypt
c, err := eg.Encrypt(m, pubK, big.NewInt(int64(15)))
if err!=nil {
	fmt.Println(err)
}

// decrypt
d, err := eg.Decrypt(c, privK)
if err!=nil {
	fmt.Println(err)
}

// check that decryption is correct
if !m.Equal(d) {
	fmt.Println("decrypted not equal to original")
}
```
