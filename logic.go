package main

import (
	"fmt"
	"os"
)

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func writeStringToFile(s string, name string) {

	file, errs := os.Create(name)
	if errs != nil {
		fmt.Println("Fallo al crear el archivo:", errs)
		return
	}
	defer file.Close()

	_, errs = file.WriteString(s)
	if errs != nil {
		fmt.Println("Falló la escritura del archivo:", errs)
		return
	}
	okMessage := fmt.Sprintf("Más detalles en el archivo: '%s'.", name)
	fmt.Println(okMessage)
}

func roundToNearestHundred(number float64) float64 {
	remainder := number - float64(int(number/100))*100
	if remainder >= 50 {
		return number + (100 - remainder)
	}
	result := number - remainder
	return result
}

func minus15(number float64) float64 {
	num := number - float64(int(number/100))*15
	return num
}
