package main

import (
	"fmt"
)

var raizInt *No[int]
var raizStr *No[string]

var strings = []string{"uva", "banana", "maçã", "laranja", "abacaxi"}
var numeros = []int{10, 5, 2, 3, 7, 6, 1, 8, 9, 4}

func main() {
	fmt.Println("Antes:", numeros)
	fmt.Println("Depois:", quicksort(numeros))

	fmt.Println("Antes:", strings)
	fmt.Println("Depois:", quicksort(strings))

	for _, v := range numeros {
		raizInt = inserir(raizInt, v)
	}

	fmt.Println("Buscar 7:", buscar(raizInt, 7))
	fmt.Println("Buscar 11:", buscar(raizInt, 11))

	for _, p := range strings {
		raizStr = inserir(raizStr, p)
	}

	fmt.Println(`Buscar "maçã":`, buscar(raizStr, "maçã"))
	fmt.Println(`Buscar "melancia":`, buscar(raizStr, "melancia"))
}
