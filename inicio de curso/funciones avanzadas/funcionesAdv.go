package main

import (
	"errors"
	"fmt"
	"strings"
)

// la siguiente funcion puede retornar dos valores
func sumar(nombre string, numeros ...int) (string, int) { //recibir cantidades indeterminadas de valores los cuales se almacenan en forma de array
	fmt.Println(numeros)
	mensaje := fmt.Sprintf("La suma de %s es: ", nombre)
	var total int
	for _, num := range numeros {
		total += num
	}
	return mensaje, total
}
func main() {
	msj, resul := sumar("victor", 4, 5, 6, 8) //
	fmt.Println(msj, resul)
	fmt.Println(factorial(5)) //factorial de 5 es 120

	//Funcion anonima no tiene nombre
	/* func() {
		fmt.Println("hola desde una funcion anonima")
	}() */
	miFunc := func(nombre string) string { //funcion anonima guardadda en una variable
		return fmt.Sprintf("hola %s desde una funcion anonima", nombre)

	}
	fmt.Println(miFunc("victor")) //si la variable no tiene parentesis solo imprimira un adireccion de memoria
	repeat := clousers(3)
	fmt.Println(repeat("hola"))
	fmt.Println(repeat("Victor"))
	main2()

	result, error := divisionErr(4, 0)
	if error == nil {
		fmt.Println("REsultado: ", result)
	} else {
		fmt.Println(error)
	}
}

// Recursividad
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	fac := n * factorial(n-1)
	return fac
}

// Clouser es una funcion que retorna otra funcion
// CLOUSURES: se debe especificar el tipo de datos que retorna y es func. tambien esta funcion retornada tiene parametros y tipos de retorno
func clousers(n int) func(cadena string) string {
	return func(cadena string) string {
		return strings.Repeat(cadena, n) //repite la cadena *n, repite la variable n veces

	}
}

//Punteros o apuntadores

func modificarValor(cadena *string) { //con asterisco le paso la direccion de memoria
	fmt.Printf("%p\n", cadena)         //imprimo referencia a memoria
	*cadena = "Hola desde una funcion" // Modifico directamente lo que contiene la direccion de memoria, algo asi como una variable global
}
func main2() {
	cadena := "Hola mundo de go"
	fmt.Printf("%p\n", &cadena) //imprimo referencia de memora
	fmt.Println("Antes de la funcion: ", cadena)
	modificarValor(&cadena)
	fmt.Println("Despues de la funcion: ", cadena)
}

// Manejo de errores,
func divisionErr(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No es posible dividir entre 0!!")
	} else {
		return dividendo / divisor, nil
	}

}
