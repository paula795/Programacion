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
	fmt.Println("ID:", p.ID, "| Nombre:", p.Nombre,
		"| Precio:", p.Precio, "| Cantidad:", p.Cantidad)
}

type Inventario struct {
	Productos map[int]Producto
}

func NuevoInventario() Inventario {
	return Inventario{make(map[int]Producto)}
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
	}
}

func main() {
	inventario := NuevoInventario()
	var opcion, id, cantidad int
	var nombre string
	var precio float64

	for opcion != 0 {
		fmt.Println("\n1.Agregar porducto 2.Precio 3.Cantidad 4.Buscar 5.Eliminar 6.Listar 0.Salir")
		fmt.Print("Opcion: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			fmt.Print("ID, nombre, precio y cantidad: ")
			fmt.Scan(&id, &nombre, &precio, &cantidad)

			if _, existe := inventario.BuscarProducto(id); existe {
				fmt.Println("Ese ID ya existe.")
			} else {
				inventario.AgregarProducto(NuevoProducto(id, nombre, precio, cantidad))
				fmt.Println("Producto agregado.")
			}

		case 2, 3:
			fmt.Print("ID del producto: ")
			fmt.Scan(&id)
			p, existe := inventario.BuscarProducto(id)

			if existe {
				if opcion == 2 {
					fmt.Print("Nuevo precio: ")
					fmt.Scan(&precio)
					p.ActualizarPrecio(precio)
				} else {
					fmt.Print("Nueva cantidad: ")
					fmt.Scan(&cantidad)
					p.ActualizarCantidad(cantidad)
				}
				inventario.AgregarProducto(p)
				fmt.Println("Producto actualizado.")
			} else {
				fmt.Println("Producto no encontrado.")
			}

		case 4:
			fmt.Print("ID del producto: ")
			fmt.Scan(&id)
			p, existe := inventario.BuscarProducto(id)
			if existe {
				p.MostrarInfo()
			} else {
				fmt.Println("Producto no encontrado.")
			}

		case 5:
			fmt.Print("ID del producto: ")
			fmt.Scan(&id)
			if _, existe := inventario.BuscarProducto(id); existe {
				inventario.EliminarProducto(id)
				fmt.Println("Producto eliminado.")
			} else {
				fmt.Println("Producto no encontrado.")
			}

		case 6:
			inventario.ListarProductos()

		case 0:
			fmt.Println("Fin del programa.")

		default:
			fmt.Println("Opcion no valida.")
		}
	}
}
