package main

import (
	"fmt"
	"unsafe"
)

// No byte padding
type nbp struct {
	a bool
	b bool
	c bool
}

// Single byte padding
type sbp struct {
	a bool 
	b int32
}

type tbp struct {
	a bool
	b int32
}

// Seven byte padding
type svnbp struct {
	a bool
	b int64
}

type np struct {
	a string
	b string
	c int32
	d int32
}

type ebp64 struct {
	a string
	b int32
	c string
	d int32
}

func main() {
	var nbp nbp
	size := unsafe.Sizeof(nbp)
	fmt.Printf("nbp   : SizeOf[%d][%p %p %p]\n", size, &nbp.a, &nbp.b, &nbp.c)
	
	var sbp sbp
	size = unsafe.Sizeof(sbp)
	fmt.Printf("sbp   : SizeOf[%d][%p %p]\n", size, &sbp.a, &sbp.b)

	var tbp tbp
	size = unsafe.Sizeof(tbp)
	fmt.Printf("sbp   : SizeOf[%d][%p %p]\n", size, &tbp.a, &tbp.b)

	var svnbp svnbp
	size = unsafe.Sizeof(svnbp)
	fmt.Printf("svnbp: SizeOf[%d][%p %p]\n", size, &svnbp.a, &svnbp.b)

	var np np
	size = unsafe.Sizeof(np)
	fmt.Printf("np   : SizeOf[%d][%p %p %p %p]\n", size, &np.a, &np.b, &np.c, &np.d)

	var ebp64 ebp64
	size = unsafe.Sizeof(ebp64)
	fmt.Printf("ebp64: SizeOf[%d][%p %p %p %p]\n", size, &ebp64.a, &ebp64.b, &ebp64.c, &ebp64.d)
}
