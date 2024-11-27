package main

import "fmt"

// Escribe una función que reciba un número n y devuelva la suma de todos los números pares desde 1 hasta n.
func ejercicio1() int {
	fmt.Println("Ingresa un numero: ")
	var n int
	var suma int

	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			suma += i
		}
	}
	return suma

}
