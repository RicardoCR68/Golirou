package main

import (
	"log"
	"sync"
)

func taylor(i int) float64 {

	sum := 0.0
	for x := i; x < 1000000000+i; x++ {

		sum += 1 / float64(x)
	}

	return sum
}

func main() {
	var wg sync.WaitGroup
	sum := 0.0
	// log.Println("Start")
	for i := 1; i < 100000000000; i += 1000000000 {

		wg.Add(1)
		i := i

		go func() {

			sum += taylor(i)
			defer wg.Done()
		}()

	}
	wg.Wait()
	log.Println(sum)
	// log.Println("Finish")
}
