package main

import "fmt"

func averageGrade(suma int, estudiantes int) int {
	promedio := suma / int(estudiantes)
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

		for nota < 0 || nota > 100 {
			fmt.Println("La nota debe estar entre 0 y 100.")
			fmt.Printf("Ingrese nuevamente la nota del estudiante %d: ", i+1)
			fmt.Scan(&nota)
		}

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
