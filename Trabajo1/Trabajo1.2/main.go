package main

import "fmt"

func main() {
	var n int

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&n)

	for i := 1; i <= 10; i++ {
		fmt.Println(n, "x", i, "=", n*i)
	}

	if n%2 == 0 {
		fmt.Println("El número es Par")
	} else {
		fmt.Println("El número es Impar")
	}

	switch {
	case n >= 1 && n <= 5:
		fmt.Println("Número pequeño")
	case n >= 6 && n <= 10:
		fmt.Println("Número mediano")
	default:
		fmt.Println("Número grande")
	}
}
