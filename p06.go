// Copyright 2011 Karl Sackett <karl.sackett@gmail.com>
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// p06.go
// Find the difference between the square of the sum of the first one hundred
// natural numbers and the sum of the squares.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sumSquares, sum int

	n, _ := strconv.Atoi(os.Args[1])
	for i := 1; i <= n; i++ {
		sum += i
		sumSquares += i * i
	}
	fmt.Println((sum * sum) - sumSquares)
}

