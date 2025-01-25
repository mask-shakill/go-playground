package main

import "fmt"

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    fmt.Println("Enter the value of n:")
    var n int
    fmt.Scanf("%d", &n)
    fmt.Printf("Fibonacci of %d is %d\n", n, fibonacci(n))
}
