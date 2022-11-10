package main

import (
	"fmt"
	"os"
)

// variables global que es accedido por tdas as funciones
var mensaje string

func funcion1() {
	mensaje = "hola desde la funcion 1"
	fmt.Println(mensaje)
}
func funcion2() {
	mensaje = "hola desde la funcion 2 "
	fmt.Println(mensaje)
}

func main() {
	mensaje = "hola desde el main"
	fmt.Println(mensaje)
	funcion1()
	funcion2()
	fmt.Println(mensaje)
	leerArchivo()
}

// defer,  la instruccion defer me indica que esa linea se ejecuta al final
func leerArchivo() {
	//cuando se nos paso un panic error podemos captar el error con un recover
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Ups! La aplicacion no finalizo de forma correcta ")
		}
	}()
	//paquete que nos permite trabajar en el sistema
	//file, error := os.Open("hola.txt")
	if file, error := os.Open("hol.txt"); error != nil {
		//fmt.Println("No se pudo leer el archivo")
		panic("No se pudo leer el archivo") //detiene el programa de manera brusca
	} else {
		defer func() {
			fmt.Println("Cerrando el archivo")
			file.Close()
		}()
		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])
		fmt.Println(contenidoArchivo)
	}
	//file.Close() // cerramos el archivo

}
