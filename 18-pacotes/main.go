package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/murilo3m/pos-go-fullcycle/tree/main/18-pacotes/matematica"
)

func main() {
	s := matematica.Soma(110, 20)
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro)
	fmt.Printf("Resultado: %v", s)
	fmt.Println(uuid.New())
}
