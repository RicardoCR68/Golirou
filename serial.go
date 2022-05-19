package main

import (
	"log"
)

func taylor() float64 {
	soma := 0.0
	div := 1.0
	for i := 1.0; i < 100000000000; i++ {
		soma = soma + (div / i)
	}

	return soma
}

func main() {
	log.Println("Inicio")
	log.Println(taylor())

}
