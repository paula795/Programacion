package main

import "fmt"

func main() {

	notas := [6][4]float64{
		{8.5, 9.0, 7.5, 8.0},
		{7.0, 8.5, 9.0, 8.0},
		{9.5, 9.0, 8.5, 9.0},
		{6.5, 7.0, 8.0, 7.5},
		{8.0, 8.5, 7.0, 9.0},
		{9.0, 8.0, 9.5, 8.5},
	}

	promedios := []float64{0, 0, 0, 0, 0, 0}
	notasAltas := []float64{0, 0, 0, 0, 0, 0}
	notasBajas := []float64{0, 0, 0, 0, 0, 0}

	sumaGeneral := 0.0

	for estudiante := 0; estudiante < 6; estudiante++ {

		suma := 0.0

		notasAltas[estudiante] = notas[estudiante][0]
		notasBajas[estudiante] = notas[estudiante][0]

		for materia := 0; materia < 4; materia++ {

			suma = suma + notas[estudiante][materia]

			if notas[estudiante][materia] > notasAltas[estudiante] {
				notasAltas[estudiante] = notas[estudiante][materia]
			}

			if notas[estudiante][materia] < notasBajas[estudiante] {
				notasBajas[estudiante] = notas[estudiante][materia]
			}
		}

		promedios[estudiante] = suma / 4

		sumaGeneral = sumaGeneral + suma
	}

	fmt.Println("Análisis de notas")
	fmt.Println("-------------------------------")

	for estudiante := 0; estudiante < 6; estudiante++ {

		fmt.Println("Estudiante", estudiante+1)
		fmt.Println("Promedio:", promedios[estudiante])
		fmt.Println("Nota más alta:", notasAltas[estudiante])
		fmt.Println("Nota más baja:", notasBajas[estudiante])
		fmt.Println("-------------------------------")
	}

	promedioGeneral := sumaGeneral / 24

	fmt.Println("Promedio general de la clase:", promedioGeneral)
}
