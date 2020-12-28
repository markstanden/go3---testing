package summy

import (
	"fmt"
	"testing"
)

func TestMySum(t *testing.T) {
	testVal := MySum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	if testVal != 55 {
		t.Errorf("mySum(1,2,3,4,5,6,7,8,9,10) expected 55, got %v", testVal)
	}
}

func ExampleMySum() {
	fmt.Println(MySum(1, 2, 3))
	// Output:
	// 6
}

func BenchmarkMySum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(MySum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, i))
	}
}
