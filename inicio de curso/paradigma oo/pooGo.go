package main

import "fmt"

/* POO SE COMPONE DE ESTOS 4 ELEMENTOS
struct=>	  Modelo(tambien llamado clases)
propiedades=> atributos
metodos=>	  comportamiento
objetos=>	  Objetos

PILARES DEL POO=>
Encapsulamiento:
Abstraccion: analiza objetos para crear clases,
Herencia:
Polimorfismo:

*/ //primero indentificar los objetos: son aquellos que tienen propiedades y conmportamientos
//ESTRUCTURAS
type Persona struct { //para crear un struct usamos la palabra reservada type
	nombre string
	edad   int
} // creo objetos desde esta estructura o clase
//CREO METODOS PARA ESTA CLASE
func (p *Persona) imprimir() { // p es una variable para aceder a sus atributos,
	//con * ago referencia a memoria a la estructura
	fmt.Printf("Nombre: %s\nEdad: %d\n", p.nombre, p.edad)
}
func (e *Persona) editNombre(nuevoNombre string) { //metodo para modivicar los valores de sus atributos
	e.nombre = nuevoNombre
}

//HERENCIA
type Empleado struct {
	Persona //con esto ya estoy implementando herencia
	sueldo  float64
}

func main() {
	per01 := Persona{"Victor", 35}
	fmt.Println(per01)
	//per01.nombre = "Moises"
	per01.editNombre("Moises")
	//fmt.Println(per01)
	per01.imprimir()
	per02 := Persona{ //si no recuerdo el orden directamente especifico con la propiedad
		nombre: "Victor",
		edad:   36}
	//fmt.Println(per02)
	per02.imprimir()
	emp01 := Empleado{
		sueldo: 500,
	} //no se puede agregar directamente nombre ni edad aqui solo los que son propios de la clase
	emp01.nombre = "pedro" // esta es la forma de asignar los atributos que se heredan
	emp01.edad = 25
	emp01.imprimir()
	fmt.Println(emp01)
}
