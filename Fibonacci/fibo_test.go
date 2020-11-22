package main

import (
	"testing"
)

func BenchmarkFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib1(10)
	}
}
func BenchmarkFib210(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib2(10)
	}
}

func BenchmarkFib310(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib3(10)
	}
}

func BenchmarkFib410(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib4(10)
	}
}

func BenchmarkFib120(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib1(20)
	}
}
func BenchmarkFib220(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib2(20)
	}
}

func BenchmarkFib320(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib3(20)
	}
}

func BenchmarkFib420(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib4(20)
	}
}
