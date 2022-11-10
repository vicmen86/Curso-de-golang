package main

import (
	"fmt"
	"strings"
)

func main() {
	if esPalindromo("Luz azul") {
		fmt.Println("La palabra es palindromo")
	} else {
		fmt.Println("La palabra no es palindromo")
	}
	esPalindromo("viaje de negocios")
}

func esPalindromo(palabra string) bool {
	fmt.Println(palabra)
	palabra = strings.ToLower(palabra) // con Replace se remplaza una letra por otra, tambien puedo remplazar espacios
	//palabra = strings.Replace(palabra, "Z", "S", 2) //remplaza la letra las veces que especifique con el unmero
	palabra = strings.ReplaceAll(palabra, " ", "") //remplazamos todo

	fmt.Println(palabra)
	fmt.Println(reverse(palabra))
	palabraInvertida := reverse(palabra)
	return palabra == palabraInvertida
}

// revertir una palabra
func reverse(cadena string) string {

	arrayCadena := strings.Split(cadena, "") //se convierte a un array y otenemos un array de letras
	arrayInvertida := make([]string, 0)
	for i := len(arrayCadena) - 1; i >= 0; i-- {
		arrayInvertida = append(arrayInvertida, arrayCadena[i])
	}

	/* 	fmt.Println(arrayCadena)
	   	fmt.Println(arrayInvertida) */
	return strings.Join(arrayInvertida, "")
}
