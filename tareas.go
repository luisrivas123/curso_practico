package main

import "fmt"

type taskList struct {
	// se define un slice de tipo task
	tasks []*task
}

// Métodos de la estructura taskList
func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

// (t *task) toma el valor que esta en la referencia
func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descricion string) {
	t.descripcion = descricion
}

func (t *task) actualizarNombren(nombre string) {
	t.nombre = nombre
}

func (t *taskList) imprimirLista() {
	for _, tarea := range t.tasks {
		fmt.Println("Nombre:", tarea.nombre)
		fmt.Println("Descripción:", tarea.descripcion)
	}
}

func (t *taskList) imprimirListaCompletados() {
	for _, tarea := range t.tasks {
		if tarea.completado {
			fmt.Println("Nombre:", tarea.nombre)
			fmt.Println("Descripción:", tarea.descripcion)
		}

	}
}

func main() {
	// Enviamos la referencia del struct
	t1 := &task{
		nombre:      "Curso de go",
		descripcion: "Curso inicial",
	}

	t2 := &task{
		nombre:      "Curso practico de go",
		descripcion: "Curso práctico",
	}

	t3 := &task{
		nombre:      "Curso practico de Python",
		descripcion: "Curso Teorico",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	// +v cuando se va imprimir una interface, para seguir el patron de propiedad -> valor
	// t := task{
	// 	nombre:      "Curso C++",
	// 	descripcion: "programación",
	// }
	// fmt.Printf("%+v\n", t)
	// t.marcarCompleta()
	// t.actualizarNombren("Curso practico")
	// t.actualizarDescripcion("Curso finalizado")
	// fmt.Printf("%+v\n", t)

	// fmt.Println(lista.tasks[0])
	lista.agregarALista(t3)
	// fmt.Println(len(lista.tasks))
	// lista.eliminarDeLista(1)

	// for i := 0; i < len(list.tasks); i++ {
	// 	fmt.Println("Index", i, "nombre", list.tasks[i].nombre)
	// }

	// for index, tarea := range list.tasks {
	// 	fmt.Println("Index", index, "nombre", tarea.nombre)
	// }

	// lista.imprimirLista()
	// lista.tasks[0].marcarCompleta()
	// fmt.Println("Tareas completadas")
	// lista.imprimirListaCompletados()

	mapaTareas := make(map[string]*taskList)

	mapaTareas["Nestor"] = lista

	t4 := &task{
		nombre:      "Curso practico de Java",
		descripcion: "Curso práctico",
	}

	t5 := &task{
		nombre:      "Curso practico de C#",
		descripcion: "Curso Teorico",
	}

	lista2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	mapaTareas["Luis"] = lista2

	fmt.Println("Tares Nestor")
	mapaTareas["Nestor"].imprimirLista()
	fmt.Println("Tares Luis")
	mapaTareas["Luis"].imprimirLista()
}
