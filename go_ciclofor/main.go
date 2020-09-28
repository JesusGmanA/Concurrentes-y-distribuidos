package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var aux, factorial int
	factorial = 1
	aux = n
	var sum float64
	sum = 0
	for i := 0; i <= n; i++ {
		for j := aux; j > 0; j-- {
			factorial *= j
		}
		sum += float64(1) / float64(factorial)
		aux--
		factorial = 1
	}
	// var sum, inicial float64
	// sum = 1 //iniciamos en 1 para contabilizar 1/0!
	// inicial = 1
	// for i := 1; i < n; i++ {
	// 	inicial /= float64(i)
	// 	sum += inicial
	// }
	fmt.Printf("Valor e: %.21f", sum)

}
