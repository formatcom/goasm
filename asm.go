package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

func asm_neg(a int) int

func neg(a int) int {
	return -1*a
}

func main() {
	fmt.Println(unsafe.Sizeof(reflect.Int))
	fmt.Println(neg(2))
	fmt.Println(asm_neg(2))
}
