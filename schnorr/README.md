## Schnorr signature
- https://en.wikipedia.org/wiki/Schnorr_signature

- [x] Hash[M || R] (where M is the msg bytes and R is a Point on the ECC, using sha256 hash function)
- [x] Generate Schnorr scheme
- [x] Sign
- [x] Verify signature


#### Usage
```go
// define new elliptic curve
ec := ecc.NewEC(big.NewInt(int64(0)), big.NewInt(int64(7)), big.NewInt(int64(11)))
// define new point
g := ecc.Point{big.NewInt(int64(11)), big.NewInt(int64(27))} // Generator
// define new random r
r := big.NewInt(int64(23))                                   // random r

// define new Schnorr crypto system using the values
schnorr, sk, err := Gen(ec, g, r)
if err!=nil {
	fmt.println(err)
}

// define message to sign
m := []byte("hola")

// also we can hash the message, but it's not mandatory, as it will be done inside the schnorr.Sign, but we can perform it now, just to check the function
h := Hash([]byte("hola"), c)
if h.String() != "34719153732582497359642109898768696927847420320548121616059449972754491425079") {
	fmt.Println("not correctly hashed")
}

s, rPoint, err := schnorr.Sign(sk, m)
if err!=nil {
	fmt.println(err)
}

// verify Schnorr signature
verified, err := Verify(schnorr.EC, sk.PubK, m, s, rPoint)
if err!=nil {
	fmt.println(err)
}
if verified {
	fmt.Println("Schnorr signature correctly verified")
}
```
