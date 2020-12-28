package main

import (
	"fmt"
	"markstanden/testing/summy"
	"testing"
)

func ExamplePrint() {
	fmt.Println(summy.MySum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	// Output:
	// 55
}

func BenchmarkPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(summy.MySum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	}
}
