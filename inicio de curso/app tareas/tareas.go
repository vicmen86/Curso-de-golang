package main

import "fmt"

//Lista de tareas app
type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) appendTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}
func (tl *TaskList) removeTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

type Task struct {
	name        string
	descripcion string
	completo    bool
}

func (t *Task) imprimir() {
	fmt.Printf("Nombre: %s\nDescripcion: %s\nCompletado: %t\n", t.name, t.descripcion, t.completo)
}

func (t *Task) marcaCompleto() {
	t.completo = true
}

func main() {
	tarea1 := Task{"Curso de go", "Completar curso este mes", false}
	tarea2 := Task{ //lo mismo que el anterior pero especificando las propiedades
		name:        "Curso de tscrip",
		descripcion: "Completar curso esta semana",
		completo:    true}
	lista := TaskList{}
	lista.appendTask(&tarea1) //agrego referencias a memoria con &
	lista.appendTask(&tarea2)
	tarea1.imprimir()
	tarea2.imprimir()

	tarea3 := Task{
		name:        "Curso de pyton",
		descripcion: "Completar curso este año",
		completo:    false}
	lista.appendTask(&tarea3)
	fmt.Println(lista)
	lista.removeTask(1)
	for i, task := range lista.tasks {
		fmt.Println(i, task.name)
	}
	principal()
	principal2()
}

// RELACIONES DE ESTRUCTURAS EN GO
//relacion uno a uno
type Usuario struct {
	nombre string
	email  string
	activo bool
}
type Estudiante struct {
	user   Usuario //variable user de tipo Usuario, es la forma de hacer la realcion uno
	codigo string
}

func principal() {
	victor := Usuario{"Victor", "viz@gmail.com", true}
	pedro := Usuario{"pedro", "pedro@gmail.com", true}
	estudianteVictor := Estudiante{user: victor, codigo: "000f9"}
	fmt.Println(pedro)
	fmt.Println(estudianteVictor)
}

//RELACIONES DE UNO A MUCHOS
type Curso struct {
	titulo string
	videos []Video //aqui especifico la relacion muchos
}

type Video struct {
	titulo string
	curso  Curso // relacion uno
}

func principal2() {
	video1 := Video{titulo: "01-Introduccion"}
	video2 := Video{titulo: "02-Introduccion"}
	curso := Curso{
		titulo: "curso de go",
		videos: []Video{video1, video2},
	}
	video1.curso = curso
	video2.curso = curso
	fmt.Println(video1.curso.titulo)
	for _, video := range curso.videos {
		fmt.Println(video.titulo)

	}
}
