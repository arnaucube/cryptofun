package schnorr

import (
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"github.com/arnaucube/cryptofun/ecc"
)

const (
	bits = 512 // 2048
)

// PubK is the public key of the Schnorr scheme
type PubK struct {
	P ecc.Point
	Q ecc.Point
}

// PrivK is the private key of the Schnorr scheme
type PrivK struct {
	PubK PubK
	A    *big.Int
}

// Schnorr is the data structure for the Schnorr scheme
type Schnorr struct {
	EC ecc.EC
	D  *big.Int // K
	G  ecc.Point
	Q  ecc.Point // P
	N  int       // order of curve
}

// Hash calculates a hash concatenating a given message bytes with a given EC Point. H(M||R)
func Hash(m []byte, c ecc.Point) *big.Int {
	var b []byte
	b = append(b, m...)
	cXBytes := c.X.Bytes()
	cYBytes := c.Y.Bytes()
	b = append(b, cXBytes...)
	b = append(b, cYBytes...)
	h := sha256.New()
	h.Write(b)
	hash := h.Sum(nil)
	r := new(big.Int).SetBytes(hash)
	return r
}

// Gen generates the Schnorr scheme
func Gen(ec ecc.EC, g ecc.Point, r *big.Int) (Schnorr, PrivK, error) {
	var err error
	var schnorr Schnorr
	var sk PrivK
	schnorr.EC = ec
	schnorr.G = g

	sk.PubK.P, _, err = ec.At(r)
	if err != nil {
		return schnorr, sk, err
	}

	orderP, err := ec.Order(sk.PubK.P)
	if err != nil {
		return schnorr, sk, err
	}

	// rand int between 1 and oerder of P
	sk.A, err = rand.Int(rand.Reader, orderP)
	if err != nil {
		return schnorr, sk, err
	}
	sk.A = big.NewInt(int64(7))
	skACopy := new(big.Int).SetBytes(sk.A.Bytes())
	// pk.Q = k x P
	sk.PubK.Q, err = ec.Mul(sk.PubK.P, skACopy)
	if err != nil {
		return schnorr, sk, err
	}
	return schnorr, sk, nil
}

// Sign performs the signature of the message m with the given private key
func (schnorr Schnorr) Sign(sk PrivK, m []byte) (*big.Int, ecc.Point, error) {
	var e *big.Int
	orderP, err := schnorr.EC.Order(sk.PubK.P)
	if err != nil {
		return e, ecc.Point{}, err
	}
	// rand k <-[1,r]
	k, err := rand.Int(rand.Reader, orderP)
	if err != nil {
		return e, ecc.Point{}, err
	}

	// R = k x P
	rPoint, err := schnorr.EC.Mul(sk.PubK.P, k)
	if err != nil {
		return e, ecc.Point{}, err
	}
	// e = H(M||R)
	e = Hash(m, rPoint)
	// a*e
	ae := new(big.Int).Mul(sk.A, e)
	// k + a*e
	kae := new(big.Int).Add(k, ae)
	// k + a*e mod r, where r is order of P
	s := new(big.Int).Mod(kae, orderP)
	return s, rPoint, nil
}

// Verify checks if the given public key matches with the given signature of the message m, in the given EC
func Verify(ec ecc.EC, pk PubK, m []byte, s *big.Int, rPoint ecc.Point) (bool, error) {
	// e = H(M||R)
	e := Hash(m, rPoint)
	eCopy := new(big.Int).SetBytes(e.Bytes())

	// e x Q
	eQ, err := ec.Mul(pk.Q, eCopy)
	if err != nil {
		return false, err
	}

	// R + e x Q
	// reQ, err := schnorr.EC.Add(rPoint, eQ)
	// if err != nil {
	// 	return false, err
	// }

	// s x P
	sp, err := ec.Mul(pk.P, s)
	if err != nil {
		return false, err
	}

	// return reQ.Equal(sp), nil
	return eQ.Equal(sp), nil
}
