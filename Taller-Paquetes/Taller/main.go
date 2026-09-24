package main

import (
	"bufio"
	"fmt"
	"os"

	"Taller/Contador_voales"
	"Taller/Conversor_monedas"
)

func main() {

	lector := bufio.NewReader(os.Stdin)

	var dolares float64
	var moneda string

	fmt.Println("Conversor de monedas")

	fmt.Print("Ingrese el valor en dólares: ")
	fmt.Fscan(lector, &dolares)

	fmt.Println("Monedas a las que es disponible convertir:")
	fmt.Println("euros")
	fmt.Println("lb")
	fmt.Println("won")
	fmt.Println("btc")

	fmt.Print("Ingrese la moneda correspondiente: ")
	fmt.Fscan(lector, &moneda)

	resultado := Conversor_monedas.Convertir(dolares, moneda)

	fmt.Println("Resultado:", resultado)

	var frase string

	fmt.Println("Contador de vocales")

	fmt.Print("Ingrese una frase: ")

	lector.ReadString('\n')
	frase, _ = lector.ReadString('\n')

	a, e, i, o, u := Contador_voales.ContarVocales(frase)

	fmt.Println("A:", a)
	fmt.Println("E:", e)
	fmt.Println("I:", i)
	fmt.Println("O:", o)
	fmt.Println("U:", u)
}
