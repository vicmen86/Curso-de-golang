package main

import (
	"fmt"
)

func main() {
	//operadores boleanos
	/* 	a := 4
	   	b := 5
	   	var res bool
	   	//operador de igualdad
	   	res = a == b
	   	fmt.Printf("%d es igual que %d? %t \n", a, b, res)
	   	//operador diferente
	   	res = a != b
	   	fmt.Printf("%d es dstinto de %d? %t \n", a, b, res)
	   	res = a < b
	   	fmt.Printf("%d es menor que %d? %t \n", a, b, res)
	   	res = a > b
	   	fmt.Printf("%d es mayor que %d? %t \n", a, b, res)
	   	//operador mayor o igual
	   	res = a >= b
	   	fmt.Printf("%d es mayor o igual que %d? %t \n", a, b, res) */

	logicaOp()
	opCondicionales()
	condicionalIf()
	opSwitch()
	bucleFor()
}

// operadores Logicos
func logicaOp() {
	//not, cambia el valor booleano
	fmt.Println(!true) // devuelve flase
	//and, ambas verdadera
	fmt.Println(true && true) //devuelve true
	//Or, almenos una verdadera
	fmt.Println(false || true) //devuelve true

}

//Condicionales op.
/* app para restuarantes
desc. de 10% hasta 100 de consumo
desc. de 20% hasta 200 de consumo
desc. de 30% en mayores a 200 de consumo
impuesto 21%
*/
func opCondicionales() {
	var consumo, descuento float64
	var datosDescuento string
	fmt.Print("ingrese consumo: ")
	fmt.Scanln(&consumo)
	if consumo >= 0 {
		if consumo <= 100 {
			datosDescuento = "10%"
			descuento = consumo * 0.10
		} else if consumo > 100 && consumo <= 200 {
			datosDescuento = "20%"
			descuento = consumo * 0.20
		} else {
			datosDescuento = "30%"
			descuento = consumo * 0.30
		}
	} else {
		fmt.Println("Error, valores ingresados no validos")
	}
	montoDescuento := consumo - descuento
	iva := montoDescuento * 0.21
	totalPagar := montoDescuento + iva

	fmt.Println("--------------FActura--------------")
	fmt.Println("Descuento que se aplica: ", datosDescuento)
	fmt.Println("Consumo: ", consumo)
	fmt.Println("Descunto: ", descuento)
	fmt.Println("Monto de descuento: ", montoDescuento)
	fmt.Println("Iva: ", iva)
	fmt.Println("Total a pagar: ", totalPagar)

}

//abreviando en en if

func condicionalIf() {
	if nombre, edad := "victo", 36; nombre == "victor" {
		fmt.Println("hola, ", nombre)
	} else {
		fmt.Printf("Nombre: %s - Edad: %d\n", nombre, edad)
	}
	//fmt.Println(nombre) no es posible por el alcance del scope

	users := make(map[string]string)
	users["victor"] = "victor@gmail.com"
	users["pedro"] = "pedro@gmail.com"
	fmt.Println(users["victor"])
	correo, ok := users["pedro"] // aqui la segunda variable tiene un valor booleano y nos dice si el valor de la primer variable existe
	fmt.Println(correo, ok)
	if correo, ok := users["juan"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("no se pudo obtener el correo")
	}
}

// switch operador de flojo que determina casos a seguir
func opSwitch() {
	var vocal string
	fmt.Print("Ingrese una letra: ")
	fmt.Scanln(&vocal)
	switch {
	case vocal == "a" || vocal == "A":
		fmt.Println(vocal, "es vocal")
	case vocal == "e" || vocal == "E":
		fmt.Println(vocal, "es vocal")
	case vocal == "i" || vocal == "I":
		fmt.Println(vocal, "es vocal")
	case vocal == "u" || vocal == "U":
		fmt.Println(vocal, "es vocal")
	case vocal == "o" || vocal == "O":
		fmt.Println(vocal, "es vocal")
	default:
		fmt.Println(vocal, "no es vocal")
	}
	// lo mismo que lo anterior pero mas simplificado
	switch vocal {
	case "a", "A", "e", "E", "i", "I", "o", "O", "u", "U":
		fmt.Println(vocal, "es vocal")
	default:
		fmt.Println(vocal, "no es vocal")
	}
}

// en go no existe bucle while, se puede remplazar por el for
func bucleFor() {
	numero := 1245
	c := 0
	for numero > 0 { //imitacion del while
		numero /= 10
		c++
	}
	fmt.Println("Cantidad de digitos", c)
	println(numero)
	// for clasico
	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	//continue y break
	for i := 0; i <= 10; i++ {
		if i == 4 {
			fmt.Println("Salto de iteracion")
			continue
		}
		if i == 6 {
			fmt.Println("detener el bucle")
			break
		}
		fmt.Println(i)
	}

	//foreach de varias maneras
	alumnos := [...]string{"juan", "carlos", "alex", "mario"}
	for i := 0; i < len(alumnos); i++ {
		fmt.Println(alumnos[i])
	}
	for indice, elemento := range alumnos { //recorre todo los valores de alumnos
		fmt.Println(indice, elemento) //podemos imprimir la variable que necesitemos, _nos permite no poner la variable
	}
}
