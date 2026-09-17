package main

import "fmt"

func main() {
	var edad int = 15
	var temperatura float32 = 21.3
	var activo bool = true
	var mensaje string = "Bieenvenid@"
	var datobyte = 255
	var dias [5]string = [5]string{"Lunes", "Martes", "Miercoles", "Juves", "Viernes"}
	var numeros []float64 = []float64{1.1, 2.2, 3.3}

	fmt.Printf("Edad: %v---TipoDato: %T\n", edad, edad)
	fmt.Printf("temperatura: %v---TipoDato: %T\n", temperatura, temperatura)
	fmt.Printf("activo: %v---TipoDato: %T\n", activo, activo)
	fmt.Printf("mensaje: %v---TipoDato: %T\n", mensaje, mensaje)
	fmt.Printf("datobyte: %v---TipoDato: %T\n", datobyte, datobyte)
	fmt.Printf("dias: %v---TipoDato: %T\n", dias, dias)
	fmt.Printf("numeros: %v---TipoDato: %T\n", numeros, numeros)

	fmt.Println("Tu edad es de: ", edad)
	fmt.Println("La temperatura es de : ", temperatura)
	fmt.Println("Tu activo es: ", true)
	fmt.Println("Tu mensaje es: ", mensaje)
	fmt.Println("Tu datobyten es: ", datobyte)
	fmt.Println("Los dias son: ", dias)
	fmt.Println("Los numeros son: ", numeros)

	var i int
	var j float64
	fmt.Println("Ingresa dos valores: ")
	fmt.Scanf("%d %f", &i, &j)

	fmt.Println("Resultado: ", (float64(i) * (j * j)))
}
