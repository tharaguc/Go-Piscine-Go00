package main

import "fmt"

var (
	s    string     = "42"
	ui   uint       = 42
	i    int        = 42
	ui8  uint8      = 42
	i16  int16      = 42
	ui32 uint32     = 42
	i64  int64      = 42
	bl   bool       = false
	f32  float32    = 42
	f64  float64    = 42
	c64  complex64  = 42 + 0i
	c128 complex128 = 42 + 21i
	ia   [1]int
	ia2  []int
	p    *int = nil
	ifc  interface{}
)

type FortyTwo struct {
}

func main() {
	ia[0] = 42
	msi := map[string]int{"42": 42}
	ft := FortyTwo{}
	ch := make(chan int)
	ch = nil
	fmt.Printf("%v is a %T.\n", s, s)
	fmt.Printf("%v is a %T.\n", ui, ui)
	fmt.Printf("%v is an %T.\n", i, i)
	fmt.Printf("%v is a %T.\n", ui8, ui8)
	fmt.Printf("%v is an %T.\n", i16, i16)
	fmt.Printf("%v is a %T.\n", ui32, ui32)
	fmt.Printf("%v is an %T.\n", i64, i64)
	fmt.Printf("%v is a %T.\n", bl, bl)
	fmt.Printf("%v is a %T.\n", f32, f32)
	fmt.Printf("%v is a %T.\n", f64, f64)
	fmt.Printf("%v is a %T.\n", c64, c64)
	fmt.Printf("%v is a %T.\n", c128, c128)
	fmt.Printf("%v is a %T.\n", ft, ft)
	fmt.Printf("%v is a %T.\n", ia, ia)
	fmt.Printf("%v is a %T.\n", msi, msi)
	fmt.Printf("%p is an %T.\n", p, p)
	fmt.Printf("%v is a %T.\n", ia2, ia2)
	fmt.Printf("%v is a %T.\n", ch, ch)
	fmt.Printf("%v is a %T.\n", ifc, ifc)
}
