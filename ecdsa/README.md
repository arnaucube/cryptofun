## ECC ECDSA
- https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm

- [x] define ECDSA data structure
- [x] ECDSA Sign
- [x] ECDSA Verify signature


#### Usage
```go
// define new elliptic curve
ec := ecc.NewEC(big.NewInt(int64(1)), big.NewInt(int64(18)), big.NewInt(int64(19)))
// define new point
g := ecc.Point{big.NewInt(int64(7)), big.NewInt(int64(11))}

// define new ECDSA system
dsa, err := NewDSA(ec, g)
if err!=nil {
	fmt.Println(err)
}

// define privK&pubK over the elliptic curve
privK := big.NewInt(int64(5))
pubK, err := dsa.PubK(privK)
if err!=nil {
	fmt.Println(err)
}

// hash value to sign
hashval := big.NewInt(int64(40))

// define r
r := big.NewInt(int64(11))

// sign hashed value
sig, err := dsa.Sign(hashval, privK, r)
if err!=nil {
	fmt.Println(err)
}

// verify signature
verified, err := dsa.Verify(hashval, sig, pubK)
if err!=nil {
	fmt.Println(err)
}
if verified {
	fmt.Println("signature correctly verified")
}
```
