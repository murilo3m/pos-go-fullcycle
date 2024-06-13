package main

type Number interface {
	~int | ~float64
}

type MyNumber int

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	m := map[string]int{"Mu": 100, "Liz": 200}
	m2 := map[string]float64{"Mu": 100.0, "Liz": 200.20}
	m3 := map[string]MyNumber{"Mu": 105, "Liz": 200}
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))

	println(Compara(10, 10.0))
}
