package main

import "fmt"

func main() {
	fmt.Println("Bienvenidos a la clase de Paquetes ")
	mensaje := saludo.Saludar("Paula")
	fmt.Println(mensaje)
}
