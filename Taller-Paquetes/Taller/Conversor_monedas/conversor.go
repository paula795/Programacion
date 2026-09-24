package Conversor_monedas

func Convertir(dolares float64, moneda string) float64 {

	var resultado float64

	switch moneda {

	case "euros":
		resultado = dolares * 0.85

	case "lb (Libras Esterlinas)":
		resultado = dolares * 0.74

	case "won (Sur Koreano)":
		resultado = dolares * 1400

	case "btc":
		resultado = dolares / 100000

	}

	return resultado
}
