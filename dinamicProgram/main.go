package main
import "fmt"
var store = map[int]int{
	1: 1,
	2: 1,
}

func fibonacciTabulate(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	previous1 := 1
	previous2 := 1
	currentVal := 0
	for i := 3; i <= n; i++ {
		// salva o segundo valor anterior
		currentVal = previous1 + previous2
		// salva o primeiro valor anterior
		previous1 = previous2
		// salva de fato o segundo em uma variavel util
		previous2 = currentVal
	}
	return currentVal
}
func fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	// garante que se o valor ja foi pré calculado não seja calculado dnv
	if store[n] != 0 {
		return store[n]
	}
	val := fibonacci(n-1) + fibonacci(n-2)
	fmt.Println(val)
	store[n] = val
	return val
}
// 0, 1, 1, 2, 3, 5, 8, 13, 21,...
func main() {
	fmt.Println(fibonacci(3))
}
