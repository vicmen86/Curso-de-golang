package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expresion string) int {
	valores := strings.Split(expresion, "+")
	var resultado int
	for _, valor := range valores { //la segunda variable de la linea 13 es un manejador de error al convertir
		num, error := strconv.Atoi(valor) // paqute  y funcion que me permite convertir de string a int
		if error != nil {
			fmt.Println(error)
			fmt.Println("Error: los numeros imgresados deben ser enteros")
			fmt.Println("El unico operador que se admite es la suma +")
		} else {
			resultado += num
		}

	}

	//fmt.Println(resultado)
	return resultado
}

func main() {
	var expresion string
	//var resultado int
	fmt.Print("=>")
	fmt.Scanln(&expresion)

	fmt.Println(sumar(expresion))
}
