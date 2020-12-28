package main

import (
	"fmt"
	"markstanden/testing/summy"
)

func main() {
	print()
}

func print() {
	fmt.Println(summy.MySum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
