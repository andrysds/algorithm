package main

import "testing"

func TestFactorial(t *testing.T) {
	want := 3628800
	got := Factorial(10)
	if got != want {
		t.Errorf("Factorial(10) = %d; want %d", got, want)
	}
}

func TestFibonacci(t *testing.T) {
	want := 55
	got := Fibonacci(10)
	if got != want {
		t.Errorf("Factorial(10) = %d; want %d", got, want)
	}
}
