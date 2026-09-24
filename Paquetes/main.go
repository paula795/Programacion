package main

import (
	"fmt"
	"practica/saludo"
)

func main() {
	fmt.Println("Bienvenidos a la clase de Paquetes ")
	mensaje := saludo.Saludar("Paula")
	fmt.Println(mensaje)
}
