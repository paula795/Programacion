package main

import "fmt"

/*
func <nombre>(param1, param2,... param n)<valores de retorno>{
	-----------------------
	-----------------------
	-----------------------
	//return en el caso de que nuestra función retorne valores
	})

*/

func saludar() {
	fmt.Println("Hola, esta es mi primera función")
}

func bienvenida(nombre string) {
	fmt.Println("Bienvenid@", nombre)
}

func sumar_restar(num1 int, num2 int) (int, int) {
	var result_sum int
	result_sum = num1 + num2
	if num2 > num1 {
		var result_rest int
		result_rest = num2 - num1
		return result_sum, result_rest
	} else {
		fmt.Println("EL resultado es 0")
		return result_sum, 0
	}

}

/* Variadic function ========== Funciones Variadicas*/

func mostrarnum(numeros ...int) {
	fmt.Println("Los numeros ingresados son: ", numeros)
}

func sumatoria(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	var usr string
	fmt.Println("Ingresa tu nombre: ")
	fmt.Scan(&usr)
	saludar()
	bienvenida(usr)

	//fmt.Println("El resultado de la suma es: ", sumar(5, 10))

	var num1, num2 int
	fmt.Println("Ingresa el primer número: ")
	fmt.Scan(&num1)
	fmt.Println("Ingresa el segundo número: ")
	fmt.Scan(&num2)
	resultadoSuma, resultadoResta := sumar_restar(num1, num2)
	fmt.Println("Los resultados son: ", resultadoSuma, resultadoResta)

	mostrarnum(5, 10, 15, 20, 25)
	fmt.Println("La sumatoria es: ", sumatoria(1, 2, 3, 4, 5, 6, 7, 8, 9))

}



-----------------------------
package main

import "fmt"

// Slices globales
var productosVendidos []string
var subtotales []float64

// Registra una venta
func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)

	productosVendidos = append(productosVendidos, nombre)
	subtotales = append(subtotales, subtotal)

	fmt.Println("Venta registrada correctamente.")
}

// Muestra las estadísticas
func MostrarEstadisticas() {
	if len(productosVendidos) == 0 {
		fmt.Println("No existen ventas registradas.")
		return
	}

	total := 0.0

	for i := 0; i < len(subtotales); i++ {
		total = total + subtotales[i]
	}

	fmt.Println("Total recaudado: $", total)
}

func main() {
	var opcion int

	for opcion != 3 {
		fmt.Println("\n--- MENÚ ---")
		fmt.Println("1. Registrar una nueva venta")
		fmt.Println("2. Mostrar estadísticas")
		fmt.Println("3. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scan(&opcion)

		switch opcion {

		case 1:
			var producto int
			var cantidad int
			var nombre string
			var precio float64

			fmt.Println("\n--- PRODUCTOS ---")
			fmt.Println("1. Arroz - $1.25")
			fmt.Println("2. Leche - $0.95")
			fmt.Println("3. Pan - $0.50")

			fmt.Print("Seleccione un producto: ")
			fmt.Scan(&producto)

			switch producto {
			case 1:
				nombre = "Arroz"
				precio = 1.25
			case 2:
				nombre = "Leche"
				precio = 0.95
			case 3:
				nombre = "Pan"
				precio = 0.50
			default:
				fmt.Println("Producto no válido.")
				continue
			}

			fmt.Print("Ingrese la cantidad vendida: ")
			fmt.Scan(&cantidad)

			RegistrarVenta(nombre, precio, cantidad)

		case 2:
			MostrarEstadisticas()

		case 3:
			fmt.Println("Programa finalizado.")

		default:
			fmt.Println("Opción no válida.")
		}
	}
}