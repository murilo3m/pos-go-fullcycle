package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "três"}
	for k, v := range numeros { // Podemos usar blank identifier com "_" na key ou value
		println(k, v)
	}

	i := 0
	for i < 10 { //Coom se fosse um "while"
		println(i)
		i++
	}

	for { // Loop infinito
		println("Hell no")
	}
}
