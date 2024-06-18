package main

import "fmt"

func fibonacci(a int, b int) int {
    return a + b
}

func main() {
    b := fibonacci(10, 20)
    fmt.Println(b)

}