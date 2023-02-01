package main

import "fmt"

func main() {
	fmt.Println("Bienvenido/a al conversor")
	fmt.Println("Selecciona una de las opciones:")
	fmt.Println("1. De metros a centímetros")
	fmt.Println("2. De centímertos a metros")

	var selection int
	fmt.Scanln(&selection)

	switch selection {
	case 1:
		fmt.Println("¿Cuántos metros?")
		var meters float32
		fmt.Scanln(&meters)
		centimeters := meters * 100
		fmt.Printf("%.2f metros son %.2f centímetros", meters, centimeters)
	case 2:
		fmt.Println("¿Cuántos centímetros?")
		var centimeters float32
		fmt.Scanln(&centimeters)
		meters := centimeters / 100
		fmt.Printf("%.2f centímetros son %.2f metros", centimeters, meters)
	default:
		fmt.Println("Selecciona una opción correcta")
	}
}
