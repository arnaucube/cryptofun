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
