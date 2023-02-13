package main

import (
	"fmt"
	"math/rand"
)

func main() {
	secretNumber := rand.Intn(100) + 1

	fmt.Println("Bienvenido/a al juego de adivinanza de números")
	fmt.Println("Adivina el número entre 1 y 100")

	var guess int
	var guesses int

	for {
		fmt.Println("Introduce tu adivinanza: ")
		fmt.Scanln(&guess)
		guesses++

		if guess < secretNumber {
			fmt.Println("Demasiado bajo")
		} else if guess > secretNumber {
			fmt.Println("Demasiado alto")
		} else {
			fmt.Printf("¡Correcto! Has necesitado %d intentos para adivinarlo", guesses)
			break
		}
	}
}
