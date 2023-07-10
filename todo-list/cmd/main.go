package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type task struct {
	name        string
	description string
	completed   bool
}

func main() {
	tasks := make(map[int]task)
	lastID := 0

	for {
		fmt.Println("Seleccione la opción deseada:")
		fmt.Println("1. Ver tareas")
		fmt.Println("2. Agregar tarea")
		fmt.Println("3. Marcar tarea como completada")
		fmt.Println("4. Eliminar tarea")
		fmt.Println("5. Salir")
		fmt.Print("Opción: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
			if len(tasks) == 0 {
				fmt.Println("No hay tareas pendientes")
			} else {
				for id, task := range tasks {
					fmt.Printf("%d. %s: %s (Completado: %t)\n", id, task.name, task.description, task.completed)
				}
			}
		case 2:
			fmt.Print("Nombre de la tarea: ")
			scanner.Scan()
			name := scanner.Text()
			fmt.Print("Descripción de la tarea: ")
			scanner.Scan()
			description := scanner.Text()

			lastID++
			tasks[lastID] = task{name: name, description: description, completed: false}
			fmt.Printf("Tarea agregada con ID %d\n", lastID)
		case 3:
			fmt.Print("Introduce el ID de la tarea completada: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())

			if task, ok := tasks[id]; ok {
				task.completed = true
				tasks[id] = task
				fmt.Printf("Tarea con ID %d completada\n", id)
			} else {
				fmt.Println("ID de tarea no válido")
			}
		case 4:
			fmt.Print("Introduce el ID de la tarea a eliminar: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())

			if _, ok := tasks[id]; ok {
				delete(tasks, id)
				fmt.Printf("Tarea con ID %d eliminada\n", id)
			} else {
				fmt.Println("ID de tarea no válido")
			}
		case 5:
			fmt.Println("Saliendo...")
			os.Exit(0)
		default:
			fmt.Println("Opción inválida")
		}

		fmt.Println(strings.Repeat("-", 30))
	}
}
