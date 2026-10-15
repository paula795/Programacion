package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func calcularTotal(producto Producto) float64 {
	return producto.Precio * float64(producto.Cantidad)
}

func tieneStock(producto Producto) bool {
	return producto.Cantidad > 0
}

func mostrarProducto(producto Producto) {
	fmt.Println("Nombre:", producto.Nombre)
	fmt.Println("Precio:", producto.Precio)
	fmt.Println("Cantidad:", producto.Cantidad)
	fmt.Println("Valor total:", calcularTotal(producto))

	if tieneStock(producto) {
		fmt.Println("Estado: Disponible")
	} else {
		fmt.Println("Estado: Sin stock")
	}

	fmt.Println("------------------")
}

func main() {
	producto1 := Producto{"Arroz", 1.25, 10}
	producto2 := Producto{"Leche", 0.95, 5}
	producto3 := Producto{"Pan", 0.50, 0}

	mostrarProducto(producto1)
	mostrarProducto(producto2)
	mostrarProducto(producto3)
}
