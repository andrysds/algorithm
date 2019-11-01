package algorithm

import "testing"

func TestFactorial(t *testing.T) {
	in := 10
	want := 3628800
	got := Factorial(in)
	if got != want {
		t.Errorf("Factorial(%d) = %d; want %d", in, got, want)
	}
}

func TestFibonacci(t *testing.T) {
	in := 10
	want := 55
	got := Fibonacci(in)
	if got != want {
		t.Errorf("Fibonacci(%d) = %d; want %d", in, got, want)
	}
}

func TestSmartFibonacci(t *testing.T) {
	in := 10
	want := 55
	got := SmartFibonacci(in)
	if got != want {
		t.Errorf("SmartFibonacci(%d) = %d; want %d", in, got, want)
	}
}
