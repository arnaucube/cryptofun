## RSA cryptosystem & Blind signature & Homomorphic Multiplication
- https://en.wikipedia.org/wiki/RSA_(cryptosystem)#
- https://en.wikipedia.org/wiki/Blind_signature
- https://en.wikipedia.org/wiki/Homomorphic_encryption

- [x] GenerateKeyPair
- [x] Encrypt
- [x] Decrypt
- [x] Blind
- [x] Blind Signature
- [x] Unblind Signature
- [x] Verify Signature
- [x] Homomorphic Multiplication


#### Usage
- Key generation, Encryption, Decryption
```go
// generate key pair
key, err := GenerateKeyPair()
if err!=nil {
	fmt.Println(err)
}
mBytes := []byte("Hi")
m := new(big.Int).SetBytes(mBytes)

// encrypt message
c := Encrypt(m, key.PubK)

// decrypt ciphertext
d := Decrypt(c, key.PrivK)
if m == d {
	fmt.Println("correctly decrypted")
}
```

- Blind signatures
```go
// key generation [Alice]
key, err := GenerateKeyPair()
if err!=nil {
	fmt.Println(err)
}

// create new message [Alice]
mBytes := []byte("Hi")
m := new(big.Int).SetBytes(mBytes)

// define r value [Alice]
rVal := big.NewInt(int64(101))

// blind message [Alice]
mBlinded := Blind(m, rVal, key.PubK)

// Blind Sign the blinded message [Bob]
sigma := BlindSign(mBlinded, key.PrivK)

// unblind the blinded signed message, and get the signature of the message [Alice]
mSigned := Unblind(sigma, rVal, key.PubK)

// verify the signature [Alice/Bob/Trudy]
verified := Verify(m, mSigned, key.PubK)
if !verified {
	fmt.Println("signature could not be verified")
}
```

- Homomorphic Multiplication
```go
// key generation [Alice]
key, err := GenerateKeyPair()
if err!=nil {
	fmt.Println(err)
}

// define values [Alice]
n1 := big.NewInt(int64(11))
n2 := big.NewInt(int64(15))

// encrypt the values [Alice]
c1 := Encrypt(n1, key.PubK)
c2 := Encrypt(n2, key.PubK)

// compute homomorphic multiplication with the encrypted values [Bob]
c3c4 := HomomorphicMul(c1, c2, key.PubK)

// decrypt the result [Alice]
d := Decrypt(c3c4, key.PrivK)

// check that the result is the expected
if !bytes.Equal(new(big.Int).Mul(n1, n2).Bytes(), d.Bytes()) {
	fmt.Println("decrypted result not equal to expected result")
}
```
