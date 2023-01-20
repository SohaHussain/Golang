package main

import (
	"fmt"
	"strconv" // for string conversion
)

// for declaring at package level use full syntax
var b int = 3

// creating a block of variables that are related or declared together
var (
	actorname    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctornumber int    = 3
	season       int    = 11
)

func main() {
	// declaring variable
	var i int
	i = 42
	fmt.Println(i)

	// declaring and initializing variable
	var x int = 2
	fmt.Println(x)

	// Go can identify datatype of variable
	a := 1
	fmt.Println(a)
	fmt.Printf("%v, %T", a, a) // v -> variable, T -> datatype
	fmt.Printf("%v, %T", b, b)

	// converting datatypes
	var c float32
	c = float32(i)
	fmt.Printf("%v, %T", c, c)

	// string conversion
	var d string
	d = strconv.Itoa(i)
	fmt.Printf("%v, %T", d, d)
}

// <---------- NOTES ----------->
// declared local variables in Go should be used
// conversions of datatypes should be explicit
// cant redeclare variables, but can shadow them
// use lowercase first letter in variable names for package scope
// use uppercase first letter to export
// there is no private scope
