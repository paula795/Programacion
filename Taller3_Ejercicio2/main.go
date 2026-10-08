package main

import "fmt"

func mayorVoto(votos map[string]int) string {

	mayor := 0
	actividadMayor := ""

	for actividad, cantidad := range votos {
		if cantidad > mayor {
			mayor = cantidad
			actividadMayor = actividad
		}
	}

	return actividadMayor
}

func main() {

	votos := map[string]int{
		"Deportes":    0,
		"Videojuegos": 0,
		"Cine":        0,
		"Musica":      0,
	}

	var voto string

	fmt.Println("Actividades")
	fmt.Println("-------------------------------")
	fmt.Println("Deportes")
	fmt.Println("Videojuegos")
	fmt.Println("Cine")
	fmt.Println("Musica")
	fmt.Println("-------------------------------")

	for numero := 1; numero <= 5; numero++ {

		fmt.Print("Ingrese el voto ", numero, ": ")
		fmt.Scan(&voto)

		votos[voto] = votos[voto] + 1
	}

	fmt.Println("-------------------------------")
	fmt.Println("Resultados")
	fmt.Println("-------------------------------")

	for actividad, cantidad := range votos {
		fmt.Println(actividad, ":", cantidad, "votos")
	}

	fmt.Println("-------------------------------")
	fmt.Println("Actividad con mayor cantidad de votos:", mayorVoto(votos))
}
