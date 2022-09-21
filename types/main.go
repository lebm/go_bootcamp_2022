package main

import "fmt"

// go is strong statically typed.
// Variables types are determined at compile time and can not be changed at runtime.
func main() {
	var a = 4
	var b = 5.8

	// Can do that. a and b are of different types
	// type of a variable can not be changed
	// a = b
	// but go allows type conversions
	a = int(b)
	fmt.Println(a, b)

	// zero values
	// numeric: 0
	// string:  ""
	// bool:    false
	// nil:     nil
	// slice:   []
	// map:     map[string]int
	// func:    func()
	// chan:    make(chan int)
	var value int
	var price float64
	var name string
	var done bool
	fmt.Println(value, price, name, done)
}
