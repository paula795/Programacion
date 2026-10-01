package main

import "fmt"

type Venta struct {
	producto string
	cantidad int
	total    float64
}

var historial []Venta

func agregarVenta(producto string, precio float64, unidades int) {
	nuevaVenta := Venta{
		producto: producto,
		cantidad: unidades,
		total:    precio * float64(unidades),
	}

	historial = append(historial, nuevaVenta)

	fmt.Printf("\nSe registró la venta de %d unidad(es) de %s.\n",
		unidades, producto)
	fmt.Printf("Valor de la venta: $%.2f\n", nuevaVenta.total)
}

func verResumen() {
	if len(historial) == 0 {
		fmt.Println("\nTodavía no se han realizado ventas.")
		return
	}

	var recaudacion float64

	fmt.Println("MENÚ")

	for _, venta := range historial {
		recaudacion += venta.total
	}

	fmt.Printf("Dinero recaudado: $%.2f\n", recaudacion)
	fmt.Printf("Ventas realizadas: %d\n", len(historial))
}

func main() {
	productos := []string{"Arroz", "Leche", "Pan"}
	valores := []float64{1.25, 0.95, 0.50}

	var seleccion int

	for seleccion != 3 {
		fmt.Println("Sistema")
		fmt.Println("1 - Realizar venta")
		fmt.Println("2 - Consultar resumen")
		fmt.Println("3 - Finalizar")
		fmt.Print("Ingrese una opción: ")
		fmt.Scan(&seleccion)

		switch seleccion {

		case 1:
			fmt.Println("\nProductos:")
			for i := 0; i < len(productos); i++ {
				fmt.Printf("%d - %s ($%.2f)\n",
					i+1, productos[i], valores[i])
			}

			var codigo int
			var unidades int

			fmt.Print("Producto: ")
			fmt.Scan(&codigo)

			if codigo < 1 || codigo > len(productos) {
				fmt.Println("El producto seleccionado no existe.")
				continue
			}

			fmt.Print("Cantidad: ")
			fmt.Scan(&unidades)

			if unidades < 1 {
				fmt.Println("Debe ingresar una cantidad válida.")
				continue
			}

			posicion := codigo - 1

			agregarVenta(
				productos[posicion],
				valores[posicion],
				unidades,
			)

		case 2:
			verResumen()

		case 3:
			fmt.Println("Gracias. El sistema ha finalizado.")

		default:
			fmt.Println("Ingrese una opción:")
		}
	}
}
