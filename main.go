package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lastDollarBlue int = 735
var oficial int = 366
var reader = bufio.NewReader(os.Stdin)

func main() {

	options := []string{"Sí", "No"}

	fmt.Printf("¿Subió el dolar? (%s): ", strings.Join(options, ", "))
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if !contains(options, choice) {
		fmt.Println("?????")
	} else {
		if choice == "Sí" {
			for {
				fmt.Println("Qué bajón. Bueno, dale.")
				fmt.Println("¿Cómo se llama lo que vas a actualizar?")
				var nameInput, _ = reader.ReadString('\n')
				nameInput = strings.TrimSpace(nameInput)

				fileNameMsg := fmt.Sprintf("Tu archivo se llamara así: '%s.txt'", nameInput)
				fmt.Println(fileNameMsg)

				fmt.Println("Precio a actualizar: ")
				pInput, _ := reader.ReadString('\n')
				pInput = strings.TrimSpace(pInput)
				p, _ := strconv.Atoi(pInput)

				precioDolar := float64(p) / float64(oficial)
				fmt.Printf("El dólar anterior valía: %d y tu precio en dólares valía: %d\n", lastDollarBlue, int(precioDolar))

				fmt.Print("El nuevo valor del dólar es: ")
				dInput, _ := reader.ReadString('\n')
				dInput = strings.TrimSpace(dInput)
				d, _ := strconv.Atoi(dInput)

				fmt.Printf("Dólar actual %d\n", d)
				resultado := precioDolar * float64(d)

				fmt.Printf("Tu precio actualizado: %d\n", int(resultado))

				resultadoMenos15 := minus15(resultado)
				precioOptimista := roundToNearestHundred(resultadoMenos15)

				text := fmt.Sprintf("Actualización de %s\nEl precio a actualizar fue: $%d\nEl dólar anterior valía: $%d \nTu precio en dólares valía: USD$%.2f\nDólar actual: $%d\nTu precio actualizado: $%.2f\nTu precio optimista (-15 por ciento): $%.f", nameInput, p, lastDollarBlue, precioDolar, d, resultado, precioOptimista)

				nameInput = fmt.Sprintf("%s.txt", nameInput)
				writeStringToFile(text, nameInput)

				fmt.Printf("¿Seguimos? (%s): ", strings.Join(options, ", "))
				lastInput, _ := reader.ReadString('\n')
				lastInput = strings.TrimSpace(lastInput)
				if lastInput == "No" {
					fmt.Println("Copado. NV")
					break
				}
			}
		} else {
			fmt.Println("Copado")
		}
	}
}
