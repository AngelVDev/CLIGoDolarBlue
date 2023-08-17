package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lastDollarBlue = 592

func main() {
	reader := bufio.NewReader(os.Stdin)
	options := []string{"Sí", "No"}

	fmt.Printf("¿Subió el dolar? (%s): ", strings.Join(options, ", "))
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if !contains(options, choice) {
		fmt.Println("?????")
	} else {
		if choice == "Sí" {
			fmt.Println("Qué bajón. Bueno, dale.")
			fmt.Print("Precio a actualizar: ")
			pInput, _ := reader.ReadString('\n')
			pInput = strings.TrimSpace(pInput)
			p, _ := strconv.Atoi(pInput)

			precioDolar := float64(p) / float64(lastDollarBlue)
			fmt.Printf("El dólar anterior valía: %d y tu precio en dólares valía: %d\n", lastDollarBlue, int(precioDolar))

			fmt.Print("El nuevo valor del dólar es: ")
			dInput, _ := reader.ReadString('\n')
			dInput = strings.TrimSpace(dInput)
			d, _ := strconv.Atoi(dInput)

			fmt.Printf("Dólar actual %d\n", d)
			resultado := precioDolar * float64(d)
			lastDollarBlue = d
			fmt.Printf("Tu precio actualizado: %d\n", int(resultado))
		} else {
			fmt.Println("Copado")
		}
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
