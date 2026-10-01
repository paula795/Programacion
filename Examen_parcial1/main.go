package main

import "fmt"

var productos = [3]string{"Arroz", "Leche", "Pan"}
var precios = [3]float64{1.25, 0.95, 0.50}

var productosVendidos = [100]string{}
var subtotales = [100]float64{}

var cantidadVentas int = 0

func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)

	productosVendidos[cantidadVentas] = nombre
	subtotales[cantidadVentas] = subtotal

	cantidadVentas++

	fmt.Println("Venta registrada correctamente.")
}

func MostrarEstadisticas() {
	if cantidadVentas == 0 {
		fmt.Println("No existen ventas registradas.")
		return 
	}
}

func main() {
	for {
		fmt.Println(" Menú ")
		fmt.Println("1. Registrar una nueva venta")
		fmt.Println("2. Mostrar estadísticas")
		fmt.Println("3. Salir")
		fmt.Print("Seleccione una opción: ")

		var opcion int
		fmt.Scan(&opcion)
		switch.opcion{
			case 1 :
			MostrarEstadisticas()
		case 3: 
			fmt.Println("Saliendo del programa")
			return
		default:
			fmt.Println("Opción no válida. Por favor, seleccione una opción válida.")
		}
	}
}
