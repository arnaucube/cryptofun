package ecc

import (
	"fmt"
	"math/big"
	"testing"
)

func TestECC(t *testing.T) {
	ec := NewEC(0, 7, 11)
	p1, p1_, err := ec.At(big.NewInt(int64(7)))
	if err != nil {
		t.Errorf(err.Error())
	}
	if !p1.Equal(Point{big.NewInt(int64(7)), big.NewInt(int64(3))}) {
		t.Errorf("p1!=(7, 11)")
	}
	if !p1_.Equal(Point{big.NewInt(int64(7)), big.NewInt(int64(8))}) {
		t.Errorf("p1_!=(7, 8)")
	}
}
func TestNeg(t *testing.T) {
	ec := NewEC(0, 7, 11)
	p1, p1_, err := ec.At(big.NewInt(int64(7)))
	if err != nil {
		t.Errorf(err.Error())
	}
	p1Neg := ec.Neg(p1)
	if !p1Neg.Equal(p1_) {
		t.Errorf("p1Neg!=p1_")
	}

}
func TestAdd(t *testing.T) {
	fmt.Println("y^2 = x^3 + 7")
	fmt.Print("ec: ")
	ec := NewEC(0, 7, 11)
	fmt.Println(ec)
	p1, _, err := ec.At(big.NewInt(int64(7)))
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Print("p1: ")
	fmt.Println(p1)
	p2, _, err := ec.At(big.NewInt(int64(6)))
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Print("p2: ")
	fmt.Println(p2)

	q, err := ec.Add(p1, p2)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Print("q: ")
	fmt.Println(q)
	if !q.Equal(Point{big.NewInt(int64(2)), big.NewInt(int64(9))}) {
		t.Errorf("q!=(2, 9)")
	}

	// check that q exists on the elliptic curve
	pt, pt_, err := ec.At(q.X)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !q.Equal(pt) && !q.Equal(pt_) {
		t.Errorf("q not exist on the elliptic curve")
	}

}
