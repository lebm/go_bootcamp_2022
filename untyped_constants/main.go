package main

import "fmt"

func main() {
	// typed constant
	const a float64 = 5.1

	// untyped constant
	const b = 6.7

	const c float64 = a * b
	const str = "Hello " + "Go!"

	const d = 5 > 10
	fmt.Println(d)

	//const x int = 5

	// x and y are of different types
	//const y float64 = 2.2 * x

	// x and y are untyped
	const x = 5
	fmt.Printf("x is %T\n", x)
	const y = 2.2 * x
	fmt.Printf("y is %T\n", y)

	const z = 2 * x
	fmt.Printf("z is %T\n", z)

	var i int = x
	var j float64 = x
	var p byte = x
	fmt.Println(i, j, p)

	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3)

	const (
		c11 = iota
		c22
		c33
	)
	fmt.Println(c11, c22, c33)

	const (
		aa = iota * 2
		bb
		cc
	)
	fmt.Println(aa, bb, cc)

	const (
		a1 = (iota * 2) + 1
		b2
		c4
	)
	fmt.Println(a1, b2, c4)

	const (
		k = -(iota + 2)
		_
		l
		m
	)
	fmt.Println(k, l, m)
}
