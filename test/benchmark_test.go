package test

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	number := 1
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}

func BenchmarkFormat(b *testing.B) {
	number := int64(1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	number := 1
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
