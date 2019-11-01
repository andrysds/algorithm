package main

// Factorial (https://en.wikipedia.org/wiki/Factorial)
func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Fibonacci (https://en.wikipedia.org/wiki/Fibonacci_number)
func Fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// SmartFibonacci use looping instead
func SmartFibonacci(n int) int {
	a, b := 0, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
