package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface { //Interface em Go é somente métodos, nao atributos
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	customer := Cliente{
		Nome:  "Murilo",
		Idade: 28,
		Ativo: true,
	}

	customer.Desativar()
	customer.Address.Cidade = "Franca"

	minhaEmpresa := Empresa{}
	Desativacao(minhaEmpresa)
}
