package matematica

type Carro struct {
	Marca string
	motor string
}

func Soma[T int | float64](a, b T) T {
	return a + b
}

/*Quando vamos "exportar" um pacote o que define se as propriedades, métodos, structs e etc
vao ser vísiveis é a letra estar maiúscula, no exemplo da Struct, a propriedade motor nao ficará visível pra quem importar esse package*/
