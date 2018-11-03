## Paillier cryptosystem & Homomorphic Addition
- https://en.wikipedia.org/wiki/Paillier_cryptosystem
- https://en.wikipedia.org/wiki/Homomorphic_encryption

- [x] GenerateKeyPair
- [x] Encrypt
- [x] Decrypt
- [x] Homomorphic Addition

#### Usage
- Encrypt, Decrypt
```go
// key generation
key, err := GenerateKeyPair()
if err!=nil {
	fmt.Println(err)
}

mBytes := []byte("Hi")
m := new(big.Int).SetBytes(mBytes)

// encryption
c := Encrypt(m, key.PubK)

// decryption
d := Decrypt(c, key.PubK, key.PrivK)
if m == d {
	fmt.Println("ciphertext decrypted correctly")
}
```

- Homomorphic Addition
```go
// key generation [Alice]
key, err := GenerateKeyPair()
if err!=nil {
	fmt.Println(err)
}

// define values [Alice]
n1 := big.NewInt(int64(110))
n2 := big.NewInt(int64(150))

// encrypt values [Alice]
c1 := Encrypt(n1, key.PubK)
c2 := Encrypt(n2, key.PubK)

// compute homomorphic addition [Bob]
c3c4 := HomomorphicAddition(c1, c2, key.PubK)

// decrypt the result [Alice]
d := Decrypt(c3c4, key.PubK, key.PrivK)
if !bytes.Equal(new(big.Int).Add(n1, n2).Bytes(), d.Bytes()) {
	fmt.Println("decrypted result not equal to expected result")
}
```
