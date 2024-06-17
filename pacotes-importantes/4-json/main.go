package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"` //Caso queira omitir o valor no json fariamos assim: `json:"-"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"Numero":2,"Saldo":200}`) //Como usamos a Go Tags (na definição da Struct) aqui parou de funcionar
	var contaX Conta
	json.Unmarshal(jsonPuro, &contaX)
	println(contaX.Saldo)

	jsonPuroDiferente := []byte(`{"n":2,"s":200}`) //Pra isso vamos usar tags
	var contaDifX Conta
	json.Unmarshal(jsonPuroDiferente, &contaDifX)
	println(contaDifX.Saldo)

}
