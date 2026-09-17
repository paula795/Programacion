package main

func averageGrade(notas []float64, estudiantes int) float64 {
	suma := 0.0

	for i := 0; i < estudiantes; i++ {
		suma = suma + notas[i]
	}

	promedio := suma / float64(estudiantes)

	return promedio
}
