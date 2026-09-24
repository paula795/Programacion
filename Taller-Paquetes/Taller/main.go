package main

import (
	"fmt"

	"Taller/Contador_voales"
	"Taller/Conversor_monedas"
)

func main() {

	var dolares float64
	var moneda string

	fmt.Println("Conversor de monedas")

	fmt.Print("Ingrese el valor en dólares: ")
	fmt.Scan(&dolares)

	fmt.Println("Monedas disponibles:")
	fmt.Println("euros")
	fmt.Println("lb")
	fmt.Println("won")
	fmt.Println("btc")

	fmt.Print("Ingrese la moneda: ")
	fmt.Scan(&moneda)

	resultado := Conversor_monedas.Convertir(dolares, moneda)

	fmt.Println("Resultado:", resultado)

	var frase string

	fmt.Println("Contador de vocales")

	fmt.Print("Ingrese una frase: ")
	fmt.Scan(&frase)

	a, e, i, o, u := Contador_voales.ContarVocales(frase)

	fmt.Println("A:", a)
	fmt.Println("E:", e)
	fmt.Println("I:", i)
	fmt.Println("O:", o)
	fmt.Println("U:", u)
}
