package main

import "fmt"

func main() {
	var palabra string
	conteo := make(map[string]int)

	fmt.Println("Ingrese la palabra o texto:")

	for i := 0; i < 25; i++ {
		fmt.Scan(&palabra)
		conteo[palabra]++
	}

	mayor := 0
	palabraMayor := ""

	fmt.Println("Conteo de palabras:")

	for palabra, cantidad := range conteo {
		fmt.Println(palabra, ":", cantidad)

		if cantidad > mayor {
			mayor = cantidad
			palabraMayor = palabra
		}
	}

	fmt.Println("La palabra que mas se repite es:", palabraMayor)
	fmt.Println("Repeticiones:", mayor)
}
