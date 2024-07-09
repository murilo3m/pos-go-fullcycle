package main

import "fmt"

func panic1() {
	panic("Panic 1")
}

func panic2() {
	panic("Panic 2")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if r == "Panic 1" {
				fmt.Println("Panic 1 recovered")
			}

			if r == "Panic 2" {
				fmt.Println("Panic 2 recovered")
			}

		}
	}()

	panic1()
}
