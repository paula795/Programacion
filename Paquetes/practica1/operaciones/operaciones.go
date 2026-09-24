package operaciones

func Suma(a, b int) int {
	return a + b
}

func sumaresta(num1, num2 int) (int, int) {
	if num1 > num2 {
		ResSuma := num1 + num2
		ResResta := num1 - num2
		return ResSuma, ResResta
	} else {
		ResSuma := num1 + num2
		ResResta := 0
		return ResSuma, ResResta
	}
}

func sumatoria(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
