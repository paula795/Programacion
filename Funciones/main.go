package main

import "fmt"

func saludar() {
	fmt.Println("Hola esta es mi primera funcion")
}

// Funcion con un parametro
func nombre(nombre string) {
	fmt.Println("Bienvenida/a", nombre)
}
func main() {
	var usr string
	fmt.Println("Ingresa tu nombre:")
	fmt.Scan(&usr)
	fmt.Println("Hola bienvenido/a:", usr)
	fmt.Println("El resultado de la suma es:", suma(4, 5))
}

// solicite dos numeros y resulva su suma
func suma(a int, b int) int {
	return a + b
}

// funcion que se llame suma resta solicitar dos numeros y retornar dos valores, resultado de la suma
// y resta
func sumaResta(a, b int) (int, int) {
	return a + b, a - b
}
