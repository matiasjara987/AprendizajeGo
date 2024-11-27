package main

import (
	"fmt"
	"strings"
)

func ejercicio2() {
	var s string = "Hola mundo"
	var invertido strings.Builder
	for x := len(s) - 1; x >= 0; x-- {
		invertido.WriteByte(s[x])

	}
	fmt.Println(invertido.String())
}

/*
Escribe una función en Go que reciba una cadena de texto y la devuelva invertida.

Ejemplo:
Entrada: "golang"
Salida: "gnalog"
*/
