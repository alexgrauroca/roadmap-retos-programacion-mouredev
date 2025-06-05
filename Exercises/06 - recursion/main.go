package main

import (
	"fmt"
	"roadmap-retos-programacion-mouredev/factorial"
	"roadmap-retos-programacion-mouredev/fibonacci"
)

/*
EJERCICIO:
Entiende el concepto de recursividad creando una función recursiva que imprima
números del 100 al 0.

DIFICULTAD EXTRA (opcional):
Utiliza el concepto de recursividad para:
  - Calcular el Factorial de un número concreto (la función recibe ese número).
  - Calcular el valor de un elemento concreto (según su posición) en la
    sucesión de Fibonacci (la función recibe la posición).
*/
func main() {
	printRecursive(100)
	fact := factorial.Factorial(12)
	fmt.Println("Factorial of 12 is", fact)
	fib := fibonacci.Fibonacci(120)
	fmt.Println("Fibonacci of 120 is", fib)
}

func printRecursive(n int) {
	if n < 0 {
		return
	}

	fmt.Println(n)
	printRecursive(n - 1)
}
