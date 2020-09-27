package main

import (
	"testing"
)

func BenchmarkFib10(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fib(10)
    }
}
func BenchmarkFib210(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fib2(10)
    }
}

func BenchmarkFib20(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fib(20)
    }
}
func BenchmarkFib220(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fib2(20)
    }
}