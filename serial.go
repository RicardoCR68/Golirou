package main

import (
	"log"
)

func taylor() float64 {

	sum := 0.0
	for i := 1; i < 100000000000; i += 1000000000 {

		for x := i; x < 1000000000+i; x++ {

			sum += 1 / float64(x)
		}
	}
	
	return sum
}


func main() {
	// log.Println("Inicio")
	log.Println(taylor())

}
