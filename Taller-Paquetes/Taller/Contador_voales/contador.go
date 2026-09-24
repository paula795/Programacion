package Contador_voales

func ContarVocales(frase string) (int, int, int, int, int) {

	var a int
	var e int
	var i int
	var o int
	var u int

	for _, letra := range frase {

		switch letra {

		case 'a', 'A':
			a++

		case 'e', 'E':
			e++

		case 'i', 'I':
			i++

		case 'o', 'O':
			o++

		case 'u', 'U':
			u++

		}
	}

	return a, e, i, o, u
}
