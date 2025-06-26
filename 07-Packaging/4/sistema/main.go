package main

import (
	"github.com/google/uuid"
	"github.com/murilo3m/pos-go-fullcycle/7-Packaging/3/math"
)

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
	println(uuid.New().String())
}
