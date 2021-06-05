package main

import (
	"fmt"
)

func main() {
	DecHexBin()
	Comps()
	Constants()
	Shifting()
	RawStringLiteral()
	Incrementing()
}

func DecHexBin() {
	for i := 0; i < 20; i++ {
		fmt.Printf("Dec: %d\t\t Hex: %#x\t Binary: %8b\n", i, i, i)
	}
}

func Comps() {
	a := (42 == 42)
	b := (42 == 43)
	c := (42 <= 42)
	d := (42 <= 41)
	e := (42 >= 42)
	f := (42 >= 43)
	g := (42 != 43)
	h := (42 != 42)
	i := (42 < 43)
	j := (42 > 43)
	fmt.Println(a, b, c, d, e, f, g, h, i, j)
}

func Constants() {
	const (
		a     = 42 // Untyped
		b int = 44 // Typed
	)
	fmt.Println(a, b)
}

func Shifting() {
	a := 42
	fmt.Printf("Dec: %d\t\tHex: %#x\tBin: %b\n", a, a, a)
	b := 42 << 1
	fmt.Printf("Dec: %d\t\tHex: %#x\tBin: %b\n", b, b, b)
}

func RawStringLiteral() {
	a := `This
	 is
	 a
	 raw
	  string
	 literal`
	fmt.Println("\n", a)
}

func Incrementing() {
	const (
		a = 2021 + iota
		b = 2021 + iota
		c = 2021 + iota
		d = 2021 + iota
	)
	fmt.Println(a, b, c, d)
}
