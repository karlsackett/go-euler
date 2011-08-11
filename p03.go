// Copyright 2011 Karl Sackett <karl.sackett@gmail.com>
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// p03.go
// Find the largest prime factor of a composite number.

package main

import (
	"big"
	"os"
	"fmt"
	"strconv"
	//"rand"
)

// Euclid's algorithm to compute the greatest common divisor of two integers.
//func euclid(a int, b int) int {
//	if b == 0 {
//		return a
//	}
//	return euclid(b, a%b)
//}

// Pollard-Rho heuristic to factor a large integer.
func pollardRho(n Int64) []Int64 {
	var f []Int64
	d := new(big.Int)
	i := 11
	//x := rand.Intn(n)
	x := new(big.Int).SetInt64(2)
	y := x
	k := 2
	for j := 0; j < 10; j++ {
		i++
		fmt.Println(x, y)
		x = ((x * x) - 1) % n
		//d = euclid((y - x), n)
		big.GcdInt(d, nil, nil, (y - x), n)
		if d != 1 && d != n {
			fmt.Println(d)
			f = append(f, d.Int64)
		}
		if i == k {
			y = x
			k = k * 2
		}
	}
	return f
}

func main() {
	var p []big.Int
	c, _ := strconv.Atoi(os.Args[1])
	p = pollardRho(c)

	fmt.Println(p)
}

