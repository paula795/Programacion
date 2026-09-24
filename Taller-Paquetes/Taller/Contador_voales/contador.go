package Contador_voales

func ContarVocales(frase string) (int, int, int, int, int) {

	var a int
	var e int
	var i int
	var o int
	var u int

	for _, letra := range frase {

		switch letra {

		case 'a', 'A', 'á', 'Á':
			a++

		case 'e', 'E', 'é', 'É':
			e++

		case 'i', 'I', 'í', 'Í':
			i++

		case 'o', 'O', 'ó', 'Ó':
			o++

		case 'u', 'U', 'ú', 'Ú', 'ü', 'Ü':
			u++
		}
	}

	return a, e, i, o, u
}
