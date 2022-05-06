package main

import (
	"log"
	"sync"
)

func main() {
	log.Println("Iniciou")

	//Canais
	c := make(chan float64)
	//

	sum := 0.

	//Grupo de espera
	wg := sync.WaitGroup{}

	//
	for i := 1; i < 1000000000; i += 1000000000 {
		//Adiciona canais de espera
		wg.Add(1)
		//

		go soma(i, c, &wg)

		sum += <-c

	}

	wg.Wait()
	log.Println(sum)

	log.Println("terminou")
}

func soma(i int, c chan float64, wg *sync.WaitGroup) {
	sum := 0.0
	for x := i; x < 1000000000+i; x++ {
		sum += 1 / float64(x)
	}

	c <- sum
	wg.Done()
}
