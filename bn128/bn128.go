package bn128

import (
	"errors"
	"math/big"
)

type Bn128 struct {
	Q             *big.Int
	R             *big.Int
	Gg1           [2]*big.Int
	Gg2           [2][2]*big.Int
	NonResidueFq2 *big.Int
	NonResidueFq6 [2]*big.Int
	Fq1           Fq
	Fq2           Fq2
	Fq6           Fq6
	Fq12          Fq12
	G1            G1
	G2            G2
}

func NewBn128() (Bn128, error) {
	var b Bn128
	q, ok := new(big.Int).SetString("21888242871839275222246405745257275088696311157297823662689037894645226208583", 10) // i
	if !ok {
		return b, errors.New("err with q")
	}
	b.Q = q
	r, ok := new(big.Int).SetString("21888242871839275222246405745257275088548364400416034343698204186575808495617", 10) // i
	if !ok {
		return b, errors.New("err with r")
	}
	b.R = r

	b.Gg1 = [2]*big.Int{
		big.NewInt(int64(1)),
		big.NewInt(int64(2)),
	}

	g2_00, ok := new(big.Int).SetString("10857046999023057135944570762232829481370756359578518086990519993285655852781", 10)
	if !ok {
		return b, errors.New("err with g2_00")
	}
	g2_01, ok := new(big.Int).SetString("11559732032986387107991004021392285783925812861821192530917403151452391805634", 10)
	if !ok {
		return b, errors.New("err with g2_00")
	}
	g2_10, ok := new(big.Int).SetString("8495653923123431417604973247489272438418190587263600148770280649306958101930", 10)
	if !ok {
		return b, errors.New("err with g2_00")
	}
	g2_11, ok := new(big.Int).SetString("4082367875863433681332203403145435568316851327593401208105741076214120093531", 10)
	if !ok {
		return b, errors.New("err with g2_00")
	}
	g2_0 := [2]*big.Int{
		g2_00,
		g2_01,
	}
	g2_1 := [2]*big.Int{
		g2_10,
		g2_11,
	}
	b.Gg2 = [2][2]*big.Int{
		g2_0,
		g2_1,
	}

	b.Fq1 = NewFq(q)
	b.NonResidueFq2, ok = new(big.Int).SetString("21888242871839275222246405745257275088696311157297823662689037894645226208582", 10) // i
	if !ok {
		return b, errors.New("err with nonResidueFq2")
	}
	b.NonResidueFq6 = [2]*big.Int{
		big.NewInt(int64(9)),
		big.NewInt(int64(1)),
	}

	b.Fq2 = Fq2{b.Fq1, b.NonResidueFq2}
	b.Fq6 = Fq6{b.Fq2, b.NonResidueFq6}
	b.Fq12 = Fq12{b.Fq6, b.Fq2, b.NonResidueFq6}

	b.G1 = NewG1(b.Fq1, b.Gg1)
	b.G2 = NewG2(b.Fq2, b.Gg2)

	return b, nil
}
