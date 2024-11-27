package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "ana"
	fmt.Println(esPalindromo(strings.ToLower(s)))

}

func esPalindromo(s string) bool {
	mid := len(s) / 2
	ult := len(s) - 1
	for i := 0; i < mid; i++ {
		if s[i] != s[ult-i] {
			return false
		}

	}
	return true
}

/*
		Descripción: Escribe una función en Go que determine si una palabra o frase es un palíndromo, es decir,
		si se lee igual de izquierda a derecha que de derecha a izquierda, ignorando los espacios y diferencias de mayúsculas/minúsculas.
	Ejemplo:
	Entrada: "A man a plan a canal Panama"
	Salida: true
	Entrada: "hello"
	Salida: false
*/
