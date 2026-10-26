package main

import "fmt"

type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{id, nombre, precio, cantidad}
}

func (p *Producto) ActualizarPrecio(precio float64) {
	p.Precio = precio
}

func (p *Producto) ActualizarCantidad(cantidad int) {
	p.Cantidad = cantidad
}

func (p Producto) MostrarInfo() {
	fmt.Println("ID:", p.ID)
	fmt.Println("Nombre:", p.Nombre)
	fmt.Println("Precio:", p.Precio)
	fmt.Println("Cantidad:", p.Cantidad)
}

type Inventario struct {
	Productos map[int]Producto
}

func (i *Inventario) AgregarProducto(p Producto) {
	i.Productos[p.ID] = p
}

func (i *Inventario) EliminarProducto(id int) {
	delete(i.Productos, id)
}

func (i Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := i.Productos[id]
	return p, existe
}

func (i Inventario) ListarProductos() {
	for _, p := range i.Productos {
		p.MostrarInfo()
		fmt.Println("-------------")
	}
}

func main() {
	inventario := Inventario{
		Productos: make(map[int]Producto),
	}

	opcion := -1

	for opcion != 0 {
		fmt.Println("\n--- INVENTARIO ---")
		fmt.Println("1. Agregar producto")
		fmt.Println("2. Buscar producto")
		fmt.Println("3. Listar productos")
		fmt.Println("4. Actualizar precio")
		fmt.Println("5. Actualizar cantidad")
		fmt.Println("6. Eliminar producto")
		fmt.Println("0. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			var id, cantidad int
			var nombre string
			var precio float64

			fmt.Print("ID: ")
			fmt.Scan(&id)
			fmt.Print("Nombre: ")
			fmt.Scan(&nombre)
			fmt.Print("Precio: ")
			fmt.Scan(&precio)
			fmt.Print("Cantidad: ")
			fmt.Scan(&cantidad)

			p := NuevoProducto(id, nombre, precio, cantidad)
			inventario.AgregarProducto(p)
			fmt.Println("Producto agregado.")

		case 2:
			var id int
			fmt.Print("Ingrese el ID: ")
			fmt.Scan(&id)

			p, existe := inventario.BuscarProducto(id)
			if existe {
				p.MostrarInfo()
			} else {
				fmt.Println("Producto no encontrado.")
			}

		case 3:
			inventario.ListarProductos()

		case 4:
			var id int
			var precio float64

			fmt.Print("ID del producto: ")
			fmt.Scan(&id)

			p, existe := inventario.BuscarProducto(id)
			if existe {
				fmt.Print("Nuevo precio: ")
				fmt.Scan(&precio)
				p.ActualizarPrecio(precio)
				inventario.AgregarProducto(p)
				fmt.Println("Precio actualizado.")
			} else {
				fmt.Println("Producto no encontrado.")
			}

		case 5:
			var id, cantidad int

			fmt.Print("ID del producto: ")
			fmt.Scan(&id)

			p, existe := inventario.BuscarProducto(id)
			if existe {
				fmt.Print("Nueva cantidad: ")
				fmt.Scan(&cantidad)
				p.ActualizarCantidad(cantidad)
				inventario.AgregarProducto(p)
				fmt.Println("Cantidad actualizada.")
			} else {
				fmt.Println("Producto no encontrado.")
			}

		case 6:
			var id int
			fmt.Print("ID del producto: ")
			fmt.Scan(&id)

			_, existe := inventario.BuscarProducto(id)
			if existe {
				inventario.EliminarProducto(id)
				fmt.Println("Producto eliminado.")
			} else {
				fmt.Println("Producto no encontrado.")
			}

		case 0:
			fmt.Println("Saliendo del inventario.")

		default:
			fmt.Println("Opción no válida.")
		}
	}
}
