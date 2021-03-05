package main

import (
	"fmt"
	"stringmatching/utils"
)

func createMask(l int) uint8 {
	var n uint8 = 1
	for i := 0; i < l; i++ {
		n |= n << 1
	}
	return n
}

func main() {
	fmt.Println("hello world!")
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	var word string = "bonjour"
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c\n", word[i])
	}
	var a string = "A"
	var z string = "Z"
	fmt.Println(a[0])
	fmt.Println(z[0])
	fmt.Println("len", len(word))
	var n uint8 = 1
	n |= 1 << len(word)
	fmt.Printf("%b\n", n)
	var h = createMask(8)
	fmt.Printf("%b\n", h)
	var m1 = utils.CreateMask(8, true)
	var m2 = utils.CreateMask(8, false)
	fmt.Printf("%08b\n", m1)
	fmt.Printf("%08b\n", m2)
}
