package ecc

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestECC(t *testing.T) {
	ec := NewEC(0, 7, 11)
	p1, p1i, err := ec.At(big.NewInt(int64(7)))
	assert.Nil(t, err)

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
	assert.Nil(t, err)

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
	assert.Nil(t, err)

	if !q.Equal(Point{big.NewInt(int64(3)), big.NewInt(int64(1))}) {
		t.Errorf("q!=(3, 1)")
	}

	// check that q exists on the elliptic curve
	pt, pti, err := ec.At(q.X)
	assert.Nil(t, err)

	if !q.Equal(pt) && !q.Equal(pti) {
		t.Errorf("q not exist on the elliptic curve")
	}

}

func TestAddSamePoint(t *testing.T) {
	ec := NewEC(0, 7, 11)
	p1 := Point{big.NewInt(int64(4)), big.NewInt(int64(7))}
	p1i := Point{big.NewInt(int64(4)), big.NewInt(int64(4))}

	q, err := ec.Add(p1, p1)
	assert.Nil(t, err)

	if !q.Equal(Point{big.NewInt(int64(6)), big.NewInt(int64(5))}) {
		t.Errorf(q.String() + " == q != (6, 5)")
	}

	q_, err := ec.Add(p1i, p1i)
	assert.Nil(t, err)

	if !q_.Equal(Point{big.NewInt(int64(6)), big.NewInt(int64(6))}) {
		t.Errorf(q_.String() + " == q_ != (6, 6)")
	}

}

func TestMulPoint1(t *testing.T) {
	ec := NewEC(0, 7, 29)
	p := Point{big.NewInt(int64(11)), big.NewInt(int64(27))}

	q, err := ec.Mul(p, big.NewInt(int64(1)))
	assert.Nil(t, err)

	if !q.Equal(Point{big.NewInt(int64(11)), big.NewInt(int64(27))}) {
		t.Errorf(q.String() + " == q != (11, 27)")
	}

	q, err = ec.Mul(p, big.NewInt(int64(2)))
	assert.Nil(t, err)

	if !q.Equal(Point{big.NewInt(int64(12)), big.NewInt(int64(13))}) {
		t.Errorf(q.String() + " == q != (12, 13)")
	}

	q, err = ec.Mul(p, big.NewInt(int64(3)))
	assert.Nil(t, err)

	if !q.Equal(Point{big.NewInt(int64(28)), big.NewInt(int64(8))}) {
		t.Errorf(q.String() + " == q != (28, 8)")
	}

	q, err = ec.Mul(p, big.NewInt(int64(4)))
	assert.Nil(t, err)

	if !q.Equal(Point{big.NewInt(int64(6)), big.NewInt(int64(22))}) {
		t.Errorf(q.String() + " == q != (6, 22)")
	}
}

func TestMulPoint2(t *testing.T) {
	ec := NewEC(0, 7, 29)
	p1 := Point{big.NewInt(int64(4)), big.NewInt(int64(19))}
	q3, err := ec.Mul(p1, big.NewInt(int64(3)))
	assert.Nil(t, err)

	if !q3.Equal(Point{big.NewInt(int64(6)), big.NewInt(int64(7))}) {
		t.Errorf(q3.String() + " == q3 != (6, 7)")
	}
	q7, err := ec.Mul(p1, big.NewInt(int64(7)))
	assert.Nil(t, err)

	if !q7.Equal(Point{big.NewInt(int64(19)), big.NewInt(int64(14))}) {
		t.Errorf(q7.String() + " == q7 != (19, 14)")
	}

	q8, err := ec.Mul(p1, big.NewInt(int64(8)))
	assert.Nil(t, err)

	if !q8.Equal(Point{big.NewInt(int64(19)), big.NewInt(int64(15))}) {
		t.Errorf(q8.String() + " == q8 != (12, 16)")
	}
}

func TestMulPoint3(t *testing.T) {
	// in this test we will multiply by a high number
	ec := NewEC(0, 7, 11)
	p := Point{big.NewInt(int64(7)), big.NewInt(int64(3))}

	q, err := ec.Mul(p, big.NewInt(int64(100)))
	assert.Nil(t, err)
	if !q.Equal(Point{big.NewInt(int64(3)), big.NewInt(int64(1))}) {
		t.Errorf(q.String() + " == q != (3, 1)")
	}

	q, err = ec.Mul(p, big.NewInt(int64(100)))
	assert.Nil(t, err)
	if !q.Equal(Point{big.NewInt(int64(3)), big.NewInt(int64(1))}) {
		t.Errorf(q.String() + " == q != (3, 1)")
	}
}

func TestMulEqualSelfAdd(t *testing.T) {
	ec := NewEC(0, 7, 29)
	p1 := Point{big.NewInt(int64(11)), big.NewInt(int64(27))}

	p1_2, err := ec.Add(p1, p1)
	assert.Nil(t, err)

	p1_3, err := ec.Add(p1_2, p1)
	assert.Nil(t, err)

	// q * 3
	q, err := ec.Mul(p1, big.NewInt(int64(3)))
	assert.Nil(t, err)

	if !q.Equal(Point{big.NewInt(int64(28)), big.NewInt(int64(8))}) {
		t.Errorf(q.String() + " == q != (28, 8)")
	}
	if !q.Equal(p1_3) {
		t.Errorf("p*3 == " + q.String() + ", p+p+p == " + p1_3.String())
	}

	// q * 4
	p1_4, err := ec.Add(p1_3, p1)
	assert.Nil(t, err)

	q, err = ec.Mul(p1, big.NewInt(int64(4)))
	assert.Nil(t, err)

	if !q.Equal(Point{big.NewInt(int64(6)), big.NewInt(int64(22))}) {
		t.Errorf(q.String() + " == q != (6, 22)")
	}
	if !q.Equal(p1_4) {
		t.Errorf("p*4 == " + q.String() + ", p+p+p+p == " + p1_4.String())
	}
}

func TestOrder(t *testing.T) {
	ec := NewEC(0, 7, 11)
	g := Point{big.NewInt(int64(7)), big.NewInt(int64(8))}
	order, err := ec.Order(g)
	assert.Nil(t, err)
	assert.Equal(t, order.Int64(), int64(12))

	// another test
	g = Point{big.NewInt(int64(2)), big.NewInt(int64(9))}
	order, err = ec.Order(g)
	assert.Nil(t, err)
	assert.Equal(t, order.Int64(), int64(4))

	// another test with another curve
	ec = NewEC(0, 7, 29)
	g = Point{big.NewInt(int64(6)), big.NewInt(int64(22))}
	order, err = ec.Order(g)
	assert.Nil(t, err)
	assert.Equal(t, order.Int64(), int64(5))

	// another test
	g = Point{big.NewInt(int64(23)), big.NewInt(int64(9))}
	order, err = ec.Order(g)
	assert.Nil(t, err)
	assert.Equal(t, order.Int64(), int64(30))
}
