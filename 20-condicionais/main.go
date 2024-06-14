package main

func main() {
	a := 1
	b := 2
	c := 3

	if a > b {
		println(a)
	} else { // Nao existe else if no go
		println(b)
	}

	if a > b && c > a { //or é ||
		println("a > b && c > a ")
	}

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	default:
		println("c")
	}
}
