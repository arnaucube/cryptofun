package ecc

import (
	"fmt"
	"math/big"
	"testing"
)

func TestECC(t *testing.T) {
	ec := NewEC(0, 7, 11)
	p1, p1i, err := ec.At(big.NewInt(int64(7)))
	if err != nil {
		t.Errorf(err.Error())
	}
	if !p1.Equal(Point{big.NewInt(int64(7)), big.NewInt(int64(3))}) {
		t.Errorf("p1!=(7, 11)")
	}
	if !p1i.Equal(Point{big.NewInt(int64(7)), big.NewInt(int64(8))}) {
		t.Errorf("p1i!=(7, 8)")
	}
}
func TestNeg(t *testing.T) {
	ec := NewEC(0, 7, 11)
	p1, p1i, err := ec.At(big.NewInt(int64(7)))
	if err != nil {
		t.Errorf(err.Error())
	}
	p1Neg := ec.Neg(p1)
	if !p1Neg.Equal(p1i) {
		t.Errorf("p1Neg!=p1i")
	}

}
func TestAdd(t *testing.T) {
	ec := NewEC(0, 7, 11)
	p1 := Point{big.NewInt(int64(4)), big.NewInt(int64(7))}
	p2 := Point{big.NewInt(int64(2)), big.NewInt(int64(2))}
	q, err := ec.Add(p1, p2)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !q.Equal(Point{big.NewInt(int64(3)), big.NewInt(int64(1))}) {
		t.Errorf("q!=(3, 1)")
	}

	// check that q exists on the elliptic curve
	pt, pti, err := ec.At(q.X)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !q.Equal(pt) && !q.Equal(pti) {
		t.Errorf("q not exist on the elliptic curve")
	}

}

func TestAddSamePoint(t *testing.T) {
	ec := NewEC(0, 7, 11)
	p1 := Point{big.NewInt(int64(4)), big.NewInt(int64(7))}
	p1i := Point{big.NewInt(int64(4)), big.NewInt(int64(4))}

	q, err := ec.Add(p1, p1)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !q.Equal(Point{big.NewInt(int64(6)), big.NewInt(int64(5))}) {
		t.Errorf("q!=(6, 5)")
	}

	q_, err := ec.Add(p1i, p1i)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !q_.Equal(Point{big.NewInt(int64(6)), big.NewInt(int64(6))}) {
		t.Errorf("q_!=(6, 6)")
	}

}

func TestMulEqualSelfAdd(t *testing.T) {
	ec := NewEC(0, 7, 29)
	p1 := Point{big.NewInt(int64(11)), big.NewInt(int64(27))}

	p1p1, err := ec.Add(p1, p1)
	if err != nil {
		t.Errorf(err.Error())
	}
	p1p1, err = ec.Add(p1p1, p1)
	if err != nil {
		t.Errorf(err.Error())
	}
	q, err := ec.Mul(p1, 3)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !q.Equal(p1p1) {
		fmt.Println(q)
		fmt.Println(p1p1)
		t.Errorf("q!=p1*p1")
	}
}

func TestMul(t *testing.T) {
	ec := NewEC(0, 7, 29)
	p1 := Point{big.NewInt(int64(4)), big.NewInt(int64(19))}
	q3, err := ec.Mul(p1, 3)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !q3.Equal(Point{big.NewInt(int64(6)), big.NewInt(int64(7))}) {
		t.Errorf("q3!=(6, 7)")
	}
	q7, err := ec.Mul(p1, 7)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !q7.Equal(Point{big.NewInt(int64(19)), big.NewInt(int64(14))}) {
		t.Errorf("q7!=(19, 14)")
	}

	q8, err := ec.Mul(p1, 8)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !q8.Equal(Point{big.NewInt(int64(19)), big.NewInt(int64(15))}) {
		t.Errorf("q8!=(19, 15)")
	}
}
