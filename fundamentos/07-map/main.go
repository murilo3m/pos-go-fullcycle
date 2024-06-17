package main

import "fmt"

func main() {
	salarios := map[string]int{"Murilo": 1000, "John": 1000000}
	//fmt.Println(salarios)

	//delete(salarios, "Murilo")

	sal := make(map[string]int)
	sal["T"] = 12

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é de %d\n", nome, salario)
	}

	// Blank identifier
	for _, salario := range salarios {
		fmt.Printf("O salario é de %d\n", salario)
	}
}
