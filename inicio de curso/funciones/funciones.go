package main

import "fmt"

func saluda(nombre string) {
	fmt.Println("hola, ", nombre)
}
func sumar(a, b int) int { //antes de las llaves debemos definior que tipo de datos retorna
	return a + b
}
func main() {
	saluda("victor")
	res := sumar(10, 20)
	fmt.Println("La suma es: ", res)
	mainreto1()
	mainreto2()
}

/*
tipos de datos: -Numeros enteros
uint, solo numeros positivos
int tanto numero positivos como negativos
byte
rune numeros muy greandes
-numeros flotantes
float de 32 y 64
complex64 y 128 numeros muy grandes
-String y booleanos
string y bool

*/

//Función principal
func mainreto1() {

	//Variables
	var a, b, suma int

	//Entrada  de datos
	fmt.Print("Ingrese N01: ")
	fmt.Scanln(&a)

	fmt.Print("Ingrese N02: ")
	fmt.Scanln(&b)

	//Llamar la función sumar
	suma = sumar(a, b)

	//Salida de datos
	fmt.Println("La suma es:", suma)
}

//función que calcula cociente
func cociente(a, b int) int {
	return a / b
}

//Función que calcula residuo
func residuo(a, b int) int {
	return a % b
}

//Función principal
func mainreto2() {

	//Variables
	var a, b, c, r int

	//Entrada  de datos
	fmt.Print("Ingrese N01: ")
	fmt.Scanln(&a)

	fmt.Print("Ingrese N02: ")
	fmt.Scanln(&b)

	//Llamar la función cociente
	c = cociente(a, b)

	//LLamar la función residuo
	r = residuo(a, b)

	//Salida de datos
	fmt.Println("Cociente:", c)
	fmt.Println("Residuo:", r)
}
