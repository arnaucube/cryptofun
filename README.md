# cryptofun [![Go Report Card](https://goreportcard.com/badge/github.com/arnaucube/cryptofun)](https://goreportcard.com/report/github.com/arnaucube/cryptofun)

Crypto algorithms from scratch. Academic purposes only.


## RSA
https://en.wikipedia.org/wiki/RSA_(cryptosystem)#
- [x] GenerateKeyPair
- [x] Encrypt
- [x] Decrypt
- [x] Blind
- [x] Blind Signature
- [x] Unblind Signature
- [x] Verify Signature
- [x] Homomorphic Multiplication

## Paillier
https://en.wikipedia.org/wiki/Paillier_cryptosystem
- [x] GenerateKeyPair
- [x] Encrypt
- [x] Decrypt
- [x] Homomorphic Addition

## Shamir Secret Sharing
https://en.wikipedia.org/wiki/Shamir%27s_Secret_Sharing
- [x] create secret sharing from number of secrets needed, number of shares, random point p, secret to share
- [x] Lagrange Interpolation to restore the secret from the shares

## Diffie-Hellman
https://en.wikipedia.org/wiki/Diffie%E2%80%93Hellman_key_exchange
- [x] key exchange

## ECC
https://en.wikipedia.org/wiki/Elliptic-curve_cryptography
- [x] define elliptic curve
- [x] get point at X
- [x] get order of a Point on the elliptic curve
- [x] Add two points on the elliptic curve
- [x] Multiply a point n times on the elliptic curve

## ECC ElGamal
https://en.wikipedia.org/wiki/ElGamal_encryption
- [x] ECC ElGamal key generation
- [x] ECC ElGamal Encrypton
- [x] ECC ElGamal Decryption

## ECC ECDSA
https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm
- [x] define ECDSA data structure
- [x] ECDSA Sign
- [x] ECDSA Verify signature


## Schnorr signature
https://en.wikipedia.org/wiki/Schnorr_signature
- [x] Hash[M || R] (where M is the msg bytes and R is a Point on the ECC, using sha256 hash function)
- [x] Generate Schnorr scheme
- [x] Sign
- [x] Verify signature


## Bn128
**[not finished]**

This is implemented followng the implementations and info from:
- https://github.com/iden3/zksnark
- https://github.com/zcash/zcash/tree/master/src/snark
- `Multiplication and Squaring on Pairing-Friendly
Fields`, Augusto Jun Devegili, Colm Ó hÉigeartaigh, Michael Scott, and Ricardo Dahab https://pdfs.semanticscholar.org/3e01/de88d7428076b2547b60072088507d881bf1.pdf
- `Optimal Pairings`, Frederik Vercauteren https://www.cosic.esat.kuleuven.be/bcrypt/optimal.pdf

- [x] Fq, Fq2, Fq6, Fq12 operations

---

To run all tests:
```
go test ./... -v
```
