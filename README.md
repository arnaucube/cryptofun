# cryptofun [![Go Report Card](https://goreportcard.com/badge/github.com/arnaucube/cryptofun)](https://goreportcard.com/report/github.com/arnaucube/cryptofun) [![Build Status](https://travis-ci.org/arnaucube/cryptofun.svg?branch=master)](https://travis-ci.org/arnaucube/cryptofun)


Crypto algorithms from scratch. Academic purposes only.



- [RSA cryptosystem & Blind signature & Homomorphic Multiplication](#rsa-cryptosystem--blind-signature--homomorphic-multiplication)
- [Paillier cryptosystem & Homomorphic Addition](#paillier-cryptosystem--homomorphic-addition)
- [Shamir Secret Sharing](#shamir-secret-sharing)
- [Diffie-Hellman](#diffie-hellman)
- [ECC](#ecc)
- [ECC ElGamal](#ecc-elgamal)
- [ECC ECDSA](#ecc-ecdsa)
- [Schnorr signature](#schnorr-signature)
- [Bn128 pairing](#bn128)
- [BLS signature](#bls)

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
Implementation of the bn128 pairing.
Code moved to https://github.com/arnaucube/go-snark/tree/master/bn128


## BLS
Boneh–Lynn–Shacham (BLS) signature scheme implemented in Go.
https://en.wikipedia.org/wiki/Boneh%E2%80%93Lynn%E2%80%93Shacham

This package uses the BN128 Go implementation from https://github.com/arnaucube/go-snark/tree/master/bn128

### Usage
```go
bls, err := NewKeys()
assert.Nil(t, err)

fmt.Println("privK:", bls.PrivK)
fmt.Println("pubK:", bls.PubK)

m := []byte("test")
sig := bls.Sign(m)
fmt.Println("signature:", sig)

verified := bls.Verify(m, sig, bls.PubK)
assert.True(t, verified)

/* out:
privK: 28151522174243194157727175362620544050084772361374505986857263387912025505082855947281432752362814690196305655335201716186584298643231993241609823412370437094839017595372164876997343464950463323765646363122343203470911131219733598647659483557447955173651057370197665461325593653581904430885385707255151472097067657072671643359241937143381958053903725229458882818033464163487351806079175441316235756460455300637131488613568714712448336232283394011955460567718918055245116200622473324300828876609569556897836255866438665750954410544846238847540023603735360532628141508114304504053826700874403280496870140784630677100277
pubK: [528167154220154970470523315181365784447502116458960328551053767278433374201 18282159022449399855128689249640771309991127595389457870089153259100566421596 19728585501269572907574045312283749798205079570296187960832716959652451330253]
signature: [[12832528436266902887734423636380781315321578271441494003771296275495461508593 6964131770814642748778827029569297554111206304527781019989920684169107205085] [6508357389516441729339280841134358160957092583050390612877406497974519092306 12073245715182483402311045895787625736998570529454024932833669602347318770866] [13520730275909614846121720877644124261162513989808465368770765804305866618385 19571107788574492009101590535904131414163790958090376021518899789800327786039]]
verified: true
*/
```


---

To run all tests:
```
go test ./... -v
```
