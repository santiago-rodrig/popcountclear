package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("expected 1 argument")
		os.Exit(1)
	}
	n, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("%d has %d bits set\n", n, PopCount(n))
	fmt.Printf("%d = %b\n", n, n)
}

func PopCount(x uint64) int {
	// x & (x-1) clears the rightmost bit of x
	bits := 0
	for x != 0 {
		newX := x & (x-1)
		if newX != x {
			bits++
		}
		x = newX
	}
	return bits
}
