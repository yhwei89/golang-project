package main

import "fmt"

import "unsafe"

func main() {
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)

	aa := "hello11111"
	cc := unsafe.Sizeof(aa)
	fmt.Println(a, b, c, d, e, f, g, h, i, aa, cc)
}
