package main

import "fmt"

func menu() {
	var opcion string

	for {
		fmt.Println()
		fmt.Println("Menú de opciones a escoger")
		fmt.Println("1. Promedio de estudiantes")
		fmt.Println("2. Suma de 1 hasta n")
		fmt.Println("3. Celsius a Fahrenheit")
		fmt.Println("4. Fahrenheit a Celsius")
		fmt.Println("0. Salir")

		fmt.Print("Elija una opción: ")
		fmt.Scan(&opcion)

		switch opcion {
		case "1":
			promedioEstudiantes()

		case "2":
			sumarNumeros()

		case "3":
			celsiusAFahrenheit()

		case "4":
			fahrenheitACelsius()

		case "0":
			fmt.Println("Programa terminado.")
			return

		case "salir":
			fmt.Println("Programa terminado.")
			return

		default:
			fmt.Println("Opción no válida.")
		}
	}
}

func averageGrade(suma int, estudiantes int) float64 {
	promedio := float64(suma) / float64(estudiantes)
	return promedio
}

func promedioEstudiantes() {
	var estudiantes int
	var nota int
	var suma int

	fmt.Print("Ingrese la cantidad de estudiantes: ")
	fmt.Scan(&estudiantes)

	for i := 0; i < estudiantes; i++ {
		fmt.Printf("Ingrese la nota del estudiante %d: ", i+1)
		fmt.Scan(&nota)

		suma = suma + nota
	}

	promedio := averageGrade(suma, estudiantes)

	fmt.Println("El promedio es:", promedio)

	if promedio >= 70 {
		fmt.Println("Curso aprobado")
	} else {
		fmt.Println("Curso reprobado")
	}

	switch {
	case promedio >= 90:
		fmt.Println("Excellent performance")
	case promedio >= 80:
		fmt.Println("Good performance")
	case promedio >= 70:
		fmt.Println("Satisfactory performance")
	default:
		fmt.Println("Needs improvement")
	}
}

func sumarNumeros() {
	var numero int
	var suma int

	fmt.Print("Ingrese el valor de n: ")
	fmt.Scan(&numero)

	for i := 1; i <= numero; i++ {
		suma = suma + i
	}

	fmt.Println("La suma es:", suma)
}

func celsiusAFahrenheit() {
	var celsius float64

	fmt.Print("Ingrese la temperatura en Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32

	fmt.Println("La temperatura en Fahrenheit es:", fahrenheit)
}

func fahrenheitACelsius() {
	var fahrenheit float64

	fmt.Print("Ingrese la temperatura en Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9

	fmt.Println("La temperatura en Celsius es:", celsius)
}

func main() {
	menu()
}
