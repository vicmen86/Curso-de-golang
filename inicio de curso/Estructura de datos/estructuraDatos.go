package main

import (
	"fmt"
)

/*
Arreglos: puedo definir la cantidad de elementos que tiene el arreglo
los arreglos no cambian su tamaño
*/
func main() {
	var numeros [5]int   //se crea un arreglo de 5 elementos
	fmt.Println(numeros) // se muestra todo cero y son valore spor defecto
	numeros[0] = 45
	numeros[2] = 45
	numeros[4] = 45
	fmt.Println(numeros)
	fmt.Println(numeros[2])
	// array con valores asinados en la misma linea
	/* nombres := [3]string{"alex", "marta", "mario"}
	fmt.Println(nombres)
	//Areglo sin especificar las dimenciones
	colores := [...]string{
		"rojo",
		"amarillo",
		"verde",
	}
	fmt.Println(colores, len(colores))
	//array con indices definidos
	monedas := [...]string{
		0: "pesos",
		2: "real",
		5: "euros",
		7: "dolares",
	}
	fmt.Println(monedas, len(monedas))
	//array nuevo desde los indices de otro array
	subMoneda := monedas[0:3] // esto aqui se convierte en un slice
	fmt.Println(subMoneda) */
	slicesDatos()
	makeDatos()
	mapasDatos()
}

// estructura de datos : Slices, son tipos de array que se puede modificar su tamaño
/*
La estructura de un slice tiene 3 datos
1 punterro al arreglo
2 longitud del arreglo al que apunta
3 capacidad
*/
func slicesDatos() {
	num := []int{1, 2, 3} //un slice no se define su tamaño
	fmt.Println(num)
	//agregar datos al array con append
	num = append(num, 4)
	num = append(num, 5)
	fmt.Println(num, len(num))
	//ejemplo de de un slice
	subSlice := num[:2]
	num[0] = 50
	fmt.Println(subSlice) // aqui subSlice tambien tiene 50 en [0] porque es un apuntador a memoria del array original

	fmt.Println(num)
	meses := []string{
		"enero",
		"abril",
		"mayo",
		"junio",
	}
	fmt.Printf("Longitud: %v, Capacidad: %v, %p \n", len(meses), cap(meses), meses)
	meses = append(meses, "julio")
	fmt.Printf("Longitud: %v, Capacidad: %v, %p \n", len(meses), cap(meses), meses)
	//la capacidad cambia y tambien la direccion de memoria porque es un nuevo arreglo

}

// estructura Make
func makeDatos() {
	// definimos un slice con una longitud y capacidad
	nume := make([]int, 3, 3) //si la longitud queda en 0 podemos agregar elementos solo con append al slice
	nume[0] = 10
	nume[1] = 20
	nume[2] = 30
	nume = append(nume, 40)
	fmt.Println(nume, len(nume), cap(nume))
}

// MAPAS:estructuras de datos con elementos desordenados
// poseen claves y valores
func mapasDatos() {
	frutas := make(map[string]int) //tipo string de las claves e int de los valores
	frutas["manzana"] = 5
	frutas["pera"] = 4
	frutas["uva"] = 9
	fmt.Println(frutas)
	// un mapa no esta limitado en cantidad de elementos
	// sin embargo tambien puedo limitar la cantidad especificando despues del int el espacio que va a tener, caso anterior 3
	//Eliminar un elemento:
	delete(frutas, "uva")
	fmt.Println(frutas)
	// si eliminamos un elemento que no esta no hace ningun cambio en el mapa
	dias := make(map[int]string) //(clave int : valor string)
	dias[1] = "domingo"
	dias[2] = "lunes"
	dias[3] = "martes"
	dias[4] = "miercoles"
	dias[5] = "jueves"
	dias[6] = "viernes"
	dias[7] = "sabado"
	fmt.Println(dias)
	// en un mapa no se puede determinar su capacidad, pero si su longitud con len()
	//mapas mas complejos que tengas otras estructuras de datos adentro
	estudiantes := make(map[string][]int)
	estudiantes["marina"] = []int{1, 2, 3, 4}
	estudiantes["mario"] = []int{5, 6, 7, 8}
	fmt.Println(estudiantes)
	fmt.Println(estudiantes["marina"])    //acedemos a clave marina
	fmt.Println(estudiantes["marina"][2]) //acedemos al array de marina a un valor especifico

}
