# cryptofun [![Go Report Card](https://goreportcard.com/badge/github.com/arnaucube/cryptofun)](https://goreportcard.com/report/github.com/arnaucube/cryptofun)

Crypto algorithms from scratch. Academic purposes only.



- [RSA cryptosystem & Blind signature & Homomorphic Multiplication](#rsa-cryptosystem--blind-signature--homomorphic-multiplication)
- [Paillier cryptosystem & Homomorphic Addition](#paillier-cryptosystem--homomorphic-addition)
- [Shamir Secret Sharing](#shamir-secret-sharing)
- [Diffie-Hellman](#diffie-hellman)
- [ECC](#ecc)
- [ECC ElGamal](#ecc-elgamal)
- [ECC ECDSA](#ecc-ecdsa)
- [Schnorr signature](#schnorr-signature)
- [Bn128](#bn128)

---

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


## Shamir Secret Sharing
- https://en.wikipedia.org/wiki/Shamir%27s_Secret_Sharing

- [x] create secret sharing from number of secrets needed, number of shares, random point p, secret to share
- [x] Lagrange Interpolation to restore the secret from the shares

#### Usage
```go
// define secret to share
k := 123456789

// define random prime
p, err := rand.Prime(rand.Reader, bits/2)
if err!=nil {
	fmt.Println(err)
}

// define how many shares want to generate
nShares := big.NewInt(int64(6))

// define how many shares are needed to recover the secret
nNeededShares := big.NewInt(int64(3))

// create the shares
shares, err := Create(
	nNeededShares,
	nShares,
	p,
	big.NewInt(int64(k)))
assert.Nil(t, err)
if err!=nil {
	fmt.Println(err)
}

// select shares to use
var sharesToUse [][]*big.Int
sharesToUse = append(sharesToUse, shares[2])
sharesToUse = append(sharesToUse, shares[1])
sharesToUse = append(sharesToUse, shares[0])

// recover the secret using Lagrange Interpolation
secr := LagrangeInterpolation(sharesToUse, p)

// check that the restored secret matches the original secret
if !bytes.Equal(k.Bytes(), secr.Bytes()) {
	fmt.Println("reconstructed secret not correspond to original secret")
}
```

## Diffie-Hellman
- https://en.wikipedia.org/wiki/Diffie%E2%80%93Hellman_key_exchange

- [x] key exchange

## ECC
- https://en.wikipedia.org/wiki/Elliptic-curve_cryptography

- [x] define elliptic curve
- [x] get point at X
- [x] get order of a Point on the elliptic curve
- [x] Add two points on the elliptic curve
- [x] Multiply a point n times on the elliptic curve

#### Usage
- ECC basic operations
```go
// define new ec
ec := NewEC(big.NewInt(int64(0)), big.NewInt(int64(7)), big.NewInt(int64(11)))

// define two points over the curve
p1 := Point{big.NewInt(int64(4)), big.NewInt(int64(7))}
p2 := Point{big.NewInt(int64(2)), big.NewInt(int64(2))}

// add the two points
q, err := ec.Add(p1, p2)
if err!=nil {
	fmt.Println(err)
}

// multiply the two points
q, err := ec.Mul(p, big.NewInt(int64(1)))
if err!=nil {
	fmt.Println(err)
}

// get order of a generator point over the elliptic curve
g := Point{big.NewInt(int64(7)), big.NewInt(int64(8))}
order, err := ec.Order(g)
if err!=nil {
	fmt.Println(err)
}
```




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



## Bn128

This is implemented followng the info and the implementations from:
- `Multiplication and Squaring on Pairing-Friendly
Fields`, Augusto Jun Devegili, Colm Ó hÉigeartaigh, Michael Scott, and Ricardo Dahab https://pdfs.semanticscholar.org/3e01/de88d7428076b2547b60072088507d881bf1.pdf
- `Optimal Pairings`, Frederik Vercauteren https://www.cosic.esat.kuleuven.be/bcrypt/optimal.pdf , https://eprint.iacr.org/2008/096.pdf
- `Double-and-Add with Relative Jacobian
Coordinates`, Björn Fay https://eprint.iacr.org/2014/1014.pdf
- `Fast and Regular Algorithms for Scalar Multiplication
over Elliptic Curves`, Matthieu Rivain https://eprint.iacr.org/2011/338.pdf
- `High-Speed Software Implementation of the Optimal Ate Pairing over Barreto–Naehrig Curves`,  Jean-Luc Beuchat, Jorge E. González-Díaz, Shigeo Mitsunari, Eiji Okamoto, Francisco Rodríguez-Henríquez, and Tadanori Teruya https://eprint.iacr.org/2010/354.pdf
- `New software speed records for cryptographic pairings`, Michael Naehrig, Ruben Niederhagen, Peter Schwabe https://cryptojedi.org/papers/dclxvi-20100714.pdf
- https://github.com/zcash/zcash/tree/master/src/snark
- https://github.com/iden3/snarkjs
- https://github.com/ethereum/py_ecc/tree/master/py_ecc/bn128

- [x] Fq, Fq2, Fq6, Fq12 operations
- [x] G1, G2 operations
- [x] preparePairing
- [x] PreComupteG1, PreComupteG2
- [x] DoubleStep, AddStep
- [x] MillerLoop
- [x] Pairing


#### Usage
First let's assume that we have these three basic functions to convert integer compositions to big integer compositions:
```go
func iToBig(a int) *big.Int {
	return big.NewInt(int64(a))
}

func iiToBig(a, b int) [2]*big.Int {
	return [2]*big.Int{iToBig(a), iToBig(b)}
}

func iiiToBig(a, b int) [2]*big.Int {
	return [2]*big.Int{iToBig(a), iToBig(b)}
}
```


- Pairing
```go
bn128, err := NewBn128()
assert.Nil(t, err)

big25 := big.NewInt(int64(25))
big30 := big.NewInt(int64(30))

g1a := bn128.G1.MulScalar(bn128.G1.G, big25)
g2a := bn128.G2.MulScalar(bn128.G2.G, big30)

g1b := bn128.G1.MulScalar(bn128.G1.G, big30)
g2b := bn128.G2.MulScalar(bn128.G2.G, big25)

pA, err := bn128.Pairing(g1a, g2a)
assert.Nil(t, err)
pB, err := bn128.Pairing(g1b, g2b)
assert.Nil(t, err)
assert.True(t, bn128.Fq12.Equal(pA, pB))
```

- Finite Fields (1, 2, 6, 12) operations
```go
// new finite field of order 1
fq1 := NewFq(iToBig(7))

// basic operations of finite field 1
res := fq1.Add(iToBig(4), iToBig(4))
res = fq1.Double(iToBig(5))
res = fq1.Sub(iToBig(5), iToBig(7))
res = fq1.Neg(iToBig(5))
res = fq1.Mul(iToBig(5), iToBig(11))
res = fq1.Inverse(iToBig(4))
res = fq1.Square(iToBig(5))

// new finite field of order 2
nonResidueFq2str := "-1" // i/j
nonResidueFq2, ok := new(big.Int).SetString(nonResidueFq2str, 10)
fq2 := Fq2{fq1, nonResidueFq2}
nonResidueFq6 := iiToBig(9, 1)

// basic operations of finite field of order 2
res := fq2.Add(iiToBig(4, 4), iiToBig(3, 4))
res = fq2.Double(iiToBig(5, 3))
res = fq2.Sub(iiToBig(5, 3), iiToBig(7, 2))
res = fq2.Neg(iiToBig(4, 4))
res = fq2.Mul(iiToBig(4, 4), iiToBig(3, 4))
res = fq2.Inverse(iiToBig(4, 4))
res = fq2.Div(iiToBig(4, 4), iiToBig(3, 4))
res = fq2.Square(iiToBig(4, 4))


// new finite field of order 6
nonResidueFq6 := iiToBig(9, 1) // TODO
fq6 := Fq6{fq2, nonResidueFq6}

// define two new values of Finite Field 6, in order to be able to perform the operations
a := [3][2]*big.Int{
	iiToBig(1, 2),
	iiToBig(3, 4),
	iiToBig(5, 6)}
b := [3][2]*big.Int{
	iiToBig(12, 11),
	iiToBig(10, 9),
	iiToBig(8, 7)}

// basic operations of finite field order 6
res := fq6.Add(a, b)
res = fq6.Sub(a, b)
res = fq6.Mul(a, b)
divRes := fq6.Div(mulRes, b)


// new finite field of order 12
q, ok := new(big.Int).SetString("21888242871839275222246405745257275088696311157297823662689037894645226208583", 10) // i
if !ok {
	fmt.Println("error parsing string to big integer")
}

fq1 := NewFq(q)
nonResidueFq2, ok := new(big.Int).SetString("21888242871839275222246405745257275088696311157297823662689037894645226208582", 10) // i
assert.True(t, ok)
nonResidueFq6 := iiToBig(9, 1)

fq2 := Fq2{fq1, nonResidueFq2}
fq6 := Fq6{fq2, nonResidueFq6}
fq12 := Fq12{fq6, fq2, nonResidueFq6}

```

- G1 operations
```go
bn128, err := NewBn128()
assert.Nil(t, err)

r1 := big.NewInt(int64(33))
r2 := big.NewInt(int64(44))

gr1 := bn128.G1.MulScalar(bn128.G1.G, bn128.Fq1.Copy(r1))
gr2 := bn128.G1.MulScalar(bn128.G1.G, bn128.Fq1.Copy(r2))

grsum1 := bn128.G1.Add(gr1, gr2)
r1r2 := bn128.Fq1.Add(r1, r2)
grsum2 := bn128.G1.MulScalar(bn128.G1.G, r1r2)

a := bn128.G1.Affine(grsum1)
b := bn128.G1.Affine(grsum2)
assert.Equal(t, a, b)
assert.Equal(t, "0x2f978c0ab89ebaa576866706b14787f360c4d6c3869efe5a72f7c3651a72ff00", utils.BytesToHex(a[0].Bytes()))
assert.Equal(t, "0x12e4ba7f0edca8b4fa668fe153aebd908d322dc26ad964d4cd314795844b62b2", utils.BytesToHex(a[1].Bytes()))
```

- G2 operations
```go
bn128, err := NewBn128()
assert.Nil(t, err)

r1 := big.NewInt(int64(33))
r2 := big.NewInt(int64(44))

gr1 := bn128.G2.MulScalar(bn128.G2.G, bn128.Fq1.Copy(r1))
gr2 := bn128.G2.MulScalar(bn128.G2.G, bn128.Fq1.Copy(r2))

grsum1 := bn128.G2.Add(gr1, gr2)
r1r2 := bn128.Fq1.Add(r1, r2)
grsum2 := bn128.G2.MulScalar(bn128.G2.G, r1r2)

a := bn128.G2.Affine(grsum1)
b := bn128.G2.Affine(grsum2)
assert.Equal(t, a, b)
```


---

To run all tests:
```
go test ./... -v
```
