package main

import (
	"crypto/sha256"
	"fmt"
)

// attribute delivered to func value is a copy of byte array, cannot change outside byte array
// must deliver the ptr of array to change outside array
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	b := [32]byte{1, 2, 3}
	zero(&b)
	fmt.Println(b[0])
}
