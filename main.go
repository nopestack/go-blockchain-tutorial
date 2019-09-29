package main

import (
	"blk/ledger"
	"fmt"
)

func main() {
	c := ledger.Ledger{}
	for x := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		hashString := fmt.Sprintf("Block ID: %v", x)
		c.Add(hashString)
	}
	b, err := c.Get("Block ID: 3")
	fmt.Printf("block [%v]: %v %v", 3, b, err)
}

