package main

import (
	"fmt"
	"practica/operaciones"
	"practica/saludo"
)

func main() {
	fmt.Println("Bienvenidos a la clase de Paquetes ")
	mensaje := saludo.Saludar("Lenin")
	a := operaciones.Suma(5, 10)
	fmt.Println(a)
	fmt.Println(mensaje)
}
