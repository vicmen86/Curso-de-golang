package main

import "fmt"

func main() {
	a := 20
	b := 10
	//suma
	result := a + b
	fmt.Println("Suma:", result)
	//resta
	result = a - b
	fmt.Println("Resta:", result)
	//multiplicacion
	result = a * b
	fmt.Println("Multiplicacion:", result)
	//division
	result = a / b
	fmt.Println("Division:", result)

	// paquete fmt funciones y herramientas
	hola := "hola"
	mundo := "mundo"
	fmt.Print(hola) // imprime sin saltos de linea
	fmt.Println(mundo)
	nombre := "victor"
	edad := 35
	fmt.Printf("hola %s y tu edad es %d \n", nombre, edad)
	fmt.Printf("hola %v y tu edad es %v \n", nombre, edad) // % para decir que lo que debe inprimir es una variable
	mensaje := fmt.Sprintf("hola %s y tu edad es %d \n", nombre, edad)
	fmt.Println("Mensaje: ", mensaje)
	//determinar el tipo de una variable
	fmt.Printf("nombre: %T \n", nombre)
	fmt.Printf("nombre: %T \n", edad)
	//ingresar datos por consola
	fmt.Print("Ingrese otro nombre: ")
	fmt.Scanln(&nombre)
	fmt.Println("Otro nombre: ", nombre)
}
