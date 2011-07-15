// Copyright 2011 Karl Sackett <karl.sackett@gmail.com>
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// p03.go
// Find the largest prime factor of a composite number.

package main

import (
	"os"
	"fmt"
	"strconv"
)

// Euclid's algorithm to compute the greatest common divisor of two integers.
func euclid(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return euclid(b, a % b)
	}
}

// Pollard-Rho heuristic to factor a large integer.
func pollard_rho(n int) []int {
	code
}

func main() {
	p []int
	c, _ := strconv.Atoi(os.Args[1])

	fmt.Println()
}

