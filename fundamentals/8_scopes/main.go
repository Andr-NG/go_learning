package main

import "fmt"

var global int

func main() {
	a := 100
	b := 200
	global = a + b // local variables overwrites global ones
	fmt.Printf("Local var a is %d.\nLocal var b is %d.\nGlobal var is %d", a, b, global)
}

