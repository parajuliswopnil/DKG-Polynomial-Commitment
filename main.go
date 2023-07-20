package main

import (
	"fmt"
)

var (
	FiniteFeild = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	g           = FiniteFeild[2]
	q           = FiniteFeild[len(FiniteFeild)-1]
	p           = FiniteFeild[len(FiniteFeild)-1] + 1

)

func main() {

	var i int64

	i = 4

	s1 := fx(i) % q

	commitment1, commitment2, commitment3 := commit()

	gpowsi := Pow(g, s1) % p

	verify := verify(commitment1, commitment2, commitment3, i)
	//  when verifying we calclulate g pow si modolus p which should be equal to verify
	if gpowsi != verify {
		panic("wrong")
	}

}

func fx(i int64) int64 {
	x := i
	fi := x*x + 5*x + 3
	fmt.Println("fi", fi)
	return fi
}

func commit() (int64, int64, int64) {
	return g, Pow(g, 5), Pow(g, 3)
}

func verify(c1, c2, c3, i int64) int64 {

	return (Pow(c1, Pow(i, 2)) * Pow(c2, Pow(i, 1)) * Pow(c3, Pow(i, 0))) % p
}

func Pow(base, exponenet int64) int64 {
	var ret int64
	var i int64
	ret = 1
	for i = 0; i < exponenet; i++ {
		ret *= base
	}
	return ret
}
