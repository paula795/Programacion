package main 

import "fmt"

func main(){
	visitas := make(map[string]int)
	visitas ["Inicio"] = 10
	visitas ["Noticias"] = 10
	visitas ["Deportes"] = 10
	visitas ["Hogar"] = 10
	fmt.Println("Las noticias de horar son: " ,visitas["Hogar"] )

	visitas["Hogar"] = 20
	fmt.Println("Actualizado Hogar ....\n las noticias de hogar son:", visitas["Hogar"])
	
	delete(visitas, "Hogar")
	fmt.Println(visitas)
	fmt.Println(visitas)
	fmt.Println("El conjunto de las visitas se clasifican en:")
	for key, val := range visitas {
	fmt.Println(key, ":", val)

	visitas["Galeria"] = 20
	fmt.Println("Las noticias de la galeria son:", visitas["Galeria"])

	visitas["Inicio"] = 200 
	fmt.Println("")


}



