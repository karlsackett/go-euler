// Copyright 2011 Karl Sackett <karl.sackett@gmail.com>
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Parts of this program are based on the Fibonacci closure
// example at the Go Playground <http://golang.org/doc/play/>.

// p02.go
// By considering the terms in the Fibonacci sequence whose values do not exceed
// four million, find the sum of the even-valued terms.

package main

import "fmt"

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func main() {
	s, v := 0, 0
	f := fib()
	for v <= 4000000 {
		v = f()
		if (v % 2 == 0) {
			s += v
		}
	}
	fmt.Println(s)
}

