// Copyright 2011 Karl Sackett <karl.sackett@gmail.com>
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// p01.go
// Find the sum of all the multiples of 3 or 5 below 1000.

package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	s := 0
	n, _ := strconv.Atoi(os.Args[1])
	for i := 0; i < n; i++ {
		if (i % 3 == 0) || (i % 5 == 0) {
			s += i
		}
	}
	fmt.Println(s)
}

