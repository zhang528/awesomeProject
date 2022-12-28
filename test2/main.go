package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

/*
数组
*/
func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	println(PopCount(c1, c2))
}

func PopCount(c1 [32]byte, c2 [32]byte) int {
	var res = 0
	for i := 0; i < 32; i++ {
		var ch1 = c1[i]
		var ch2 = c2[i]
		for i := 0; i < 8; i++ {
			if (ch1>>i)&1 != (ch2>>i)&1 {
				res++
			}
		}
	}
	return res
}
func encrypt(s string, b int64) string {
	var h hash.Hash
	switch b {
	case 384:
		h = sha512.New384()
	case 512:
		h = sha512.New()
	default:
		h = sha256.New()
	}
	h.Write([]byte(s))
	return string(h.Sum(nil))
}
