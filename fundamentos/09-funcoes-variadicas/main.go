package main

import (
	"fmt"
)

func main() {
	valor := sum(1, 2, 3, 4, 5)

	fmt.Println(valor)
}

// Usar por exemplo para uma funcao de configuracao de algum bd ou algo do tipo
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
